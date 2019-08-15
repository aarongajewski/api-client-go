// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1IncidentsIncidentIDTeamAssignments Assign a team to an incident for the incident
// swagger:model postV1IncidentsIncidentIdTeamAssignments
type PostV1IncidentsIncidentIDTeamAssignments struct {

	// The team ID to associate to the incident
	// Required: true
	TeamID *string `json:"team_id"`
}

// Validate validates this post v1 incidents incident Id team assignments
func (m *PostV1IncidentsIncidentIDTeamAssignments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeamID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentsIncidentIDTeamAssignments) validateTeamID(formats strfmt.Registry) error {

	if err := validate.Required("team_id", "body", m.TeamID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentsIncidentIDTeamAssignments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentsIncidentIDTeamAssignments) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentsIncidentIDTeamAssignments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
