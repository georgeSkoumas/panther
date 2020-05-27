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

// Rule rule
//
// swagger:model Rule
type Rule struct {

	// body
	// Required: true
	Body Body `json:"body"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt ModifyTime `json:"createdAt"`

	// created by
	// Required: true
	CreatedBy UserID `json:"createdBy"`

	// dedup period minutes
	// Required: true
	DedupPeriodMinutes DedupPeriodMinutes `json:"dedupPeriodMinutes"`

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

	// log types
	// Required: true
	LogTypes TypeSet `json:"logTypes"`

	// reference
	// Required: true
	Reference Reference `json:"reference"`

	// reports
	// Required: true
	Reports Reports `json:"reports"`

	// runbook
	// Required: true
	Runbook Runbook `json:"runbook"`

	// severity
	// Required: true
	Severity Severity `json:"severity"`

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

// Validate validates this rule
func (m *Rule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDedupPeriodMinutes(formats); err != nil {
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

	if err := m.validateLogTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunbook(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
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

func (m *Rule) validateBody(formats strfmt.Registry) error {

	if err := m.Body.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body")
		}
		return err
	}

	return nil
}

func (m *Rule) validateCreatedAt(formats strfmt.Registry) error {

	if err := m.CreatedAt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("createdAt")
		}
		return err
	}

	return nil
}

func (m *Rule) validateCreatedBy(formats strfmt.Registry) error {

	if err := m.CreatedBy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("createdBy")
		}
		return err
	}

	return nil
}

func (m *Rule) validateDedupPeriodMinutes(formats strfmt.Registry) error {

	if err := m.DedupPeriodMinutes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dedupPeriodMinutes")
		}
		return err
	}

	return nil
}

func (m *Rule) validateDescription(formats strfmt.Registry) error {

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("description")
		}
		return err
	}

	return nil
}

func (m *Rule) validateDisplayName(formats strfmt.Registry) error {

	if err := m.DisplayName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("displayName")
		}
		return err
	}

	return nil
}

func (m *Rule) validateEnabled(formats strfmt.Registry) error {

	if err := m.Enabled.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("enabled")
		}
		return err
	}

	return nil
}

func (m *Rule) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *Rule) validateLastModified(formats strfmt.Registry) error {

	if err := m.LastModified.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastModified")
		}
		return err
	}

	return nil
}

func (m *Rule) validateLastModifiedBy(formats strfmt.Registry) error {

	if err := m.LastModifiedBy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastModifiedBy")
		}
		return err
	}

	return nil
}

func (m *Rule) validateLogTypes(formats strfmt.Registry) error {

	if err := validate.Required("logTypes", "body", m.LogTypes); err != nil {
		return err
	}

	if err := m.LogTypes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("logTypes")
		}
		return err
	}

	return nil
}

func (m *Rule) validateReference(formats strfmt.Registry) error {

	if err := m.Reference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reference")
		}
		return err
	}

	return nil
}

func (m *Rule) validateReports(formats strfmt.Registry) error {

	if err := m.Reports.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reports")
		}
		return err
	}

	return nil
}

func (m *Rule) validateRunbook(formats strfmt.Registry) error {

	if err := m.Runbook.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("runbook")
		}
		return err
	}

	return nil
}

func (m *Rule) validateSeverity(formats strfmt.Registry) error {

	if err := m.Severity.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("severity")
		}
		return err
	}

	return nil
}

func (m *Rule) validateTags(formats strfmt.Registry) error {

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

func (m *Rule) validateTests(formats strfmt.Registry) error {

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

func (m *Rule) validateVersionID(formats strfmt.Registry) error {

	if err := m.VersionID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("versionId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rule) UnmarshalBinary(b []byte) error {
	var res Rule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
