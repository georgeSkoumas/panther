// Code generated by go-swagger; DO NOT EDIT.

package models

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

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Policy policy
//
// swagger:model Policy
type Policy struct {

	// auto remediation Id
	// Required: true
	AutoRemediationID AutoRemediationID `json:"autoRemediationId"`

	// auto remediation parameters
	// Required: true
	AutoRemediationParameters AutoRemediationParameters `json:"autoRemediationParameters"`

	// body
	// Required: true
	Body Body `json:"body"`

	// compliance status
	// Required: true
	ComplianceStatus ComplianceStatus `json:"complianceStatus"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt ModifyTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy UserID `json:"createdBy"`

	// description
	// Required: true
	Description Description `json:"description"`

	// display name
	// Required: true
	DisplayName DisplayName `json:"displayName"`

	// enabled
	// Required: true
	Enabled Enabled `json:"enabled"`

	// id
	// Required: true
	ID ID `json:"id"`

	// last modified
	// Required: true
	// Format: date-time
	LastModified ModifyTime `json:"lastModified"`

	// last modified by
	// Required: true
	LastModifiedBy UserID `json:"lastModifiedBy"`

	// reference
	// Required: true
	Reference Reference `json:"reference"`

	// reports
	// Required: true
	Reports Reports `json:"reports"`

	// resource types
	// Required: true
	ResourceTypes TypeSet `json:"resourceTypes"`

	// runbook
	// Required: true
	Runbook Runbook `json:"runbook"`

	// severity
	// Required: true
	Severity Severity `json:"severity"`

	// suppressions
	// Required: true
	Suppressions Suppressions `json:"suppressions"`

	// tags
	// Required: true
	Tags Tags `json:"tags"`

	// tests
	// Required: true
	Tests TestSuite `json:"tests"`

	// version Id
	// Required: true
	VersionID VersionID `json:"versionId"`
}

// Validate validates this policy
func (m *Policy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoRemediationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoRemediationParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComplianceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunbook(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuppressions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Policy) validateAutoRemediationID(formats strfmt.Registry) error {

	if err := m.AutoRemediationID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("autoRemediationId")
		}
		return err
	}

	return nil
}

func (m *Policy) validateAutoRemediationParameters(formats strfmt.Registry) error {

	if err := m.AutoRemediationParameters.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("autoRemediationParameters")
		}
		return err
	}

	return nil
}

func (m *Policy) validateBody(formats strfmt.Registry) error {

	if err := m.Body.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body")
		}
		return err
	}

	return nil
}

func (m *Policy) validateComplianceStatus(formats strfmt.Registry) error {

	if err := m.ComplianceStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("complianceStatus")
		}
		return err
	}

	return nil
}

func (m *Policy) validateCreatedAt(formats strfmt.Registry) error {

	if err := m.CreatedAt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("createdAt")
		}
		return err
	}

	return nil
}

func (m *Policy) validateCreatedBy(formats strfmt.Registry) error {

	if err := m.CreatedBy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("createdBy")
		}
		return err
	}

	return nil
}

func (m *Policy) validateDescription(formats strfmt.Registry) error {

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *Policy) validateDisplayName(formats strfmt.Registry) error {

	if err := m.DisplayName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("displayName")
		}
		return err
	}

	return nil
}

func (m *Policy) validateEnabled(formats strfmt.Registry) error {

	if err := m.Enabled.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("enabled")
		}
		return err
	}

	return nil
}

func (m *Policy) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Policy) validateLastModified(formats strfmt.Registry) error {

	if err := m.LastModified.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastModified")
		}
		return err
	}

	return nil
}

func (m *Policy) validateLastModifiedBy(formats strfmt.Registry) error {

	if err := m.LastModifiedBy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastModifiedBy")
		}
		return err
	}

	return nil
}

func (m *Policy) validateReference(formats strfmt.Registry) error {

	if err := m.Reference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reference")
		}
		return err
	}

	return nil
}

func (m *Policy) validateReports(formats strfmt.Registry) error {

	if err := m.Reports.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reports")
		}
		return err
	}

	return nil
}

func (m *Policy) validateResourceTypes(formats strfmt.Registry) error {

	if err := validate.Required("resourceTypes", "body", m.ResourceTypes); err != nil {
		return err
	}

	if err := m.ResourceTypes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("resourceTypes")
		}
		return err
	}

	return nil
}

func (m *Policy) validateRunbook(formats strfmt.Registry) error {

	if err := m.Runbook.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("runbook")
		}
		return err
	}

	return nil
}

func (m *Policy) validateSeverity(formats strfmt.Registry) error {

	if err := m.Severity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("severity")
		}
		return err
	}

	return nil
}

func (m *Policy) validateSuppressions(formats strfmt.Registry) error {

	if err := validate.Required("suppressions", "body", m.Suppressions); err != nil {
		return err
	}

	if err := m.Suppressions.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("suppressions")
		}
		return err
	}

	return nil
}

func (m *Policy) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	if err := m.Tags.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tags")
		}
		return err
	}

	return nil
}

func (m *Policy) validateTests(formats strfmt.Registry) error {

	if err := validate.Required("tests", "body", m.Tests); err != nil {
		return err
	}

	if err := m.Tests.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tests")
		}
		return err
	}

	return nil
}

func (m *Policy) validateVersionID(formats strfmt.Registry) error {

	if err := m.VersionID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("versionId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Policy) UnmarshalBinary(b []byte) error {
	var res Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
