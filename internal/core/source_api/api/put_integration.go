package api

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/panther-labs/panther/api/lambda/source/models"
	pollermodels "github.com/panther-labs/panther/internal/compliance/snapshot_poller/models/poller"
	awspoller "github.com/panther-labs/panther/internal/compliance/snapshot_poller/pollers/aws"
	"github.com/panther-labs/panther/pkg/awsbatch/sqsbatch"
	"github.com/panther-labs/panther/pkg/genericapi"
)

var (
	putIntegrationInternalError = &genericapi.InternalError{Message: "Failed to add source. Please try again later"}
)

// PutIntegration adds a set of new integrations in a batch.
func (api API) PutIntegration(input *models.PutIntegrationInput) (*models.SourceIntegration, error) {
	// Validate the new integration
	reason, passing, err := evaluateIntegrationFunc(api, &models.CheckIntegrationInput{
		AWSAccountID:      input.AWSAccountID,
		IntegrationType:   input.IntegrationType,
		IntegrationLabel:  input.IntegrationLabel,
		EnableCWESetup:    input.CWEEnabled,
		EnableRemediation: input.RemediationEnabled,
		S3Bucket:          input.S3Bucket,
		S3Prefix:          input.S3Prefix,
		KmsKey:            input.KmsKey,
	})
	if err != nil {
		return nil, putIntegrationInternalError
	}
	if !passing {
		zap.L().Warn("PutIntegration: resource has a misconfiguration",
			zap.Error(err),
			zap.String("reason", reason),
			zap.Any("input", input))
		return nil, &genericapi.InvalidInputError{
			Message: fmt.Sprintf("source %s did not pass configuration check because of %s",
				*input.IntegrationLabel, reason),
		}
	}

	// Filter out existing integrations
	if err := api.integrationAlreadyExists(input); err != nil {
		return nil, err
	}

	// Get ready to add appropriate permissions to the SQS queue
	permissionAdded := false
	defer func() {
		if err != nil {
			zap.L().Error("failed to put integration", zap.Error(err))
			// In case there has been any error, try to undo granting of permissions to SQS queue.
			if permissionAdded {
				if undoErr := DisableExternalSnsTopicSubscription(*input.AWSAccountID); undoErr != nil {
					zap.L().Error("failed to remove SQS permission for integration. SQS queue has additional permissions that have to be removed manually",
						zap.Error(undoErr),
						zap.Error(err))
				}
			}
		}
	}()

	switch aws.StringValue(input.IntegrationType) {
	case models.IntegrationTypeAWS3:
		permissionAdded, err = AllowExternalSnsTopicSubscription(*input.AWSAccountID)
		if err != nil {
			zap.L().Error("Failed to add permissions to log processor queue", zap.Error(errors.WithStack(err)))
			return nil, putIntegrationInternalError
		}
		err = addGlueTables(input.LogTypes)
		if err != nil {
			zap.L().Error("Failed to add glue tables to glue catalog", zap.Error(errors.WithStack(err)))
			return nil, putIntegrationInternalError
		}
	}

	// Generate the new integration
	newIntegration := generateNewIntegration(input)

	// Write to DynamoDB
	if err = dynamoClient.PutItem(integrationToItem(newIntegration)); err != nil {
		err = errors.Wrap(err, "Failed to store source integration in DDB")
		return nil, putIntegrationInternalError
	}

	// Return early to skip sending to the snapshot queue
	if aws.BoolValue(input.SkipScanQueue) {
		return newIntegration, nil
	}

	if *input.IntegrationType == models.IntegrationTypeAWSScan {
		err = api.FullScan(&models.FullScanInput{Integrations: []*models.SourceIntegrationMetadata{&newIntegration.SourceIntegrationMetadata}})
		if err != nil {
			err = errors.Wrap(err, "failed to trigger scanning of resources")
			return nil, putIntegrationInternalError
		}
	}
	return newIntegration, nil
}

func (api API) integrationAlreadyExists(input *models.PutIntegrationInput) error {
	// avoid inserting if already done
	existingIntegrations, err := api.ListIntegrations(&models.ListIntegrationsInput{})
	if err != nil {
		zap.L().Error("failed to fetch integrations", zap.Error(errors.WithStack(err)))
		return putIntegrationInternalError
	}

	for _, existingIntegration := range existingIntegrations {
		if *existingIntegration.IntegrationType == *input.IntegrationType {
			switch *existingIntegration.IntegrationType {
			case models.IntegrationTypeAWSScan:
				if *existingIntegration.AWSAccountID == *input.AWSAccountID {
					// We can only have one cloudsec integration for each account
					return &genericapi.InvalidInputError{
						Message: fmt.Sprintf("Source account %s already onboarded", *input.AWSAccountID),
					}
				}
				return nil
			case models.IntegrationTypeAWS3:
				if *existingIntegration.AWSAccountID == *input.AWSAccountID &&
					*existingIntegration.IntegrationLabel == *input.IntegrationLabel {
					// Log sources for same account need to have different labels, need to update if the log types have changed
					return &genericapi.InvalidInputError{
						Message: fmt.Sprintf("Log source for account %s with label %s already onboarded",
							*input.AWSAccountID,
							*input.IntegrationLabel),
					}
				}
			}
		}
	}

	return nil
}

// FullScan schedules scans for each Resource type for each integration.
//
// Each Resource type is sent within its own SQS message.
func (api API) FullScan(input *models.FullScanInput) error {
	var sqsEntries []*sqs.SendMessageBatchRequestEntry

	// For each integration, add a ScanMsg to the queue per service
	for _, integration := range input.Integrations {
		for resourceType := range awspoller.ServicePollers {
			scanMsg := &pollermodels.ScanMsg{
				Entries: []*pollermodels.ScanEntry{
					{
						AWSAccountID:  integration.AWSAccountID,
						IntegrationID: integration.IntegrationID,
						ResourceType:  aws.String(resourceType),
					},
				},
			}

			messageBodyBytes, err := jsoniter.MarshalToString(scanMsg)
			if err != nil {
				return &genericapi.InternalError{Message: err.Error()}
			}

			sqsEntries = append(sqsEntries, &sqs.SendMessageBatchRequestEntry{
				// Generates an ID of: IntegrationID-AWSResourceType
				Id: aws.String(
					*integration.IntegrationID + "-" + strings.Replace(resourceType, ".", "", -1),
				),
				MessageBody: aws.String(messageBodyBytes),
			})
		}
	}

	zap.L().Info(
		"scheduling new scans",
		zap.String("queueUrl", env.SnapshotPollersQueueURL),
		zap.Int("count", len(sqsEntries)),
	)

	// Batch send all the messages to SQS
	_, err := sqsbatch.SendMessageBatch(sqsClient, maxElapsedTime, &sqs.SendMessageBatchInput{
		Entries:  sqsEntries,
		QueueUrl: &env.SnapshotPollersQueueURL,
	})
	return err
}

func generateNewIntegration(input *models.PutIntegrationInput) *models.SourceIntegration {
	metadata := models.SourceIntegrationMetadata{
		CreatedAtTime:    aws.Time(time.Now()),
		CreatedBy:        input.UserID,
		IntegrationID:    aws.String(uuid.New().String()),
		IntegrationLabel: input.IntegrationLabel,
		IntegrationType:  input.IntegrationType,
	}

	switch aws.StringValue(input.IntegrationType) {
	case models.IntegrationTypeAWSScan:
		metadata.AWSAccountID = input.AWSAccountID
		metadata.CWEEnabled = input.CWEEnabled
		metadata.RemediationEnabled = input.RemediationEnabled
		metadata.ScanIntervalMins = input.ScanIntervalMins
		metadata.StackName = aws.String(getStackName(*input.IntegrationType, *input.IntegrationLabel))
	case models.IntegrationTypeAWS3:
		metadata.AWSAccountID = input.AWSAccountID
		metadata.S3Bucket = input.S3Bucket
		metadata.S3Prefix = input.S3Prefix
		metadata.KmsKey = input.KmsKey
		metadata.LogTypes = input.LogTypes
		metadata.StackName = aws.String(getStackName(*input.IntegrationType, *input.IntegrationLabel))
		metadata.LogProcessingRole = aws.String(generateLogProcessingRoleArn(*input.AWSAccountID, *input.IntegrationLabel))
	}
	return &models.SourceIntegration{
		SourceIntegrationMetadata: metadata,
	}
}
