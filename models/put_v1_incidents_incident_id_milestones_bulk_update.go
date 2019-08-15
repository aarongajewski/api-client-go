// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PutV1IncidentsIncidentIDMilestonesBulkUpdate Update a list of milestones on an incident
// swagger:model putV1IncidentsIncidentIdMilestonesBulkUpdate
type PutV1IncidentsIncidentIDMilestonesBulkUpdate struct {

	// bulk
	// Enum: [true]
	Bulk string `json:"bulk,omitempty"`

	// milestones
	// Required: true
	Milestones []*PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestonesItems0 `json:"milestones"`
}

// Validate validates this put v1 incidents incident Id milestones bulk update
func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBulk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMilestones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var putV1IncidentsIncidentIdMilestonesBulkUpdateTypeBulkPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["true"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		putV1IncidentsIncidentIdMilestonesBulkUpdateTypeBulkPropEnum = append(putV1IncidentsIncidentIdMilestonesBulkUpdateTypeBulkPropEnum, v)
	}
}

const (

	// PutV1IncidentsIncidentIDMilestonesBulkUpdateBulkTrue captures enum value "true"
	PutV1IncidentsIncidentIDMilestonesBulkUpdateBulkTrue string = "true"
)

// prop value enum
func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdate) validateBulkEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, putV1IncidentsIncidentIdMilestonesBulkUpdateTypeBulkPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdate) validateBulk(formats strfmt.Registry) error {

	if swag.IsZero(m.Bulk) { // not required
		return nil
	}

	// value enum
	if err := m.validateBulkEnum("bulk", "body", m.Bulk); err != nil {
		return err
	}

	return nil
}

func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdate) validateMilestones(formats strfmt.Registry) error {

	if err := validate.Required("milestones", "body", m.Milestones); err != nil {
		return err
	}

	for i := 0; i < len(m.Milestones); i++ {
		if swag.IsZero(m.Milestones[i]) { // not required
			continue
		}

		if m.Milestones[i] != nil {
			if err := m.Milestones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("milestones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdate) UnmarshalBinary(b []byte) error {
	var res PutV1IncidentsIncidentIDMilestonesBulkUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestonesItems0 put v1 incidents incident ID milestones bulk update milestones items0
// swagger:model PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestonesItems0
type PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestonesItems0 struct {

	// occurred at
	// Required: true
	// Format: date-time
	OccurredAt *strfmt.DateTime `json:"occurred_at"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this put v1 incidents incident ID milestones bulk update milestones items0
func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestonesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOccurredAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestonesItems0) validateOccurredAt(formats strfmt.Registry) error {

	if err := validate.Required("occurred_at", "body", m.OccurredAt); err != nil {
		return err
	}

	if err := validate.FormatOf("occurred_at", "body", "date-time", m.OccurredAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestonesItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestonesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestonesItems0) UnmarshalBinary(b []byte) error {
	var res PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestonesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
