// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PongEntity Simple endpoint to verify your API connection is working
// swagger:model PongEntity
type PongEntity struct {

	// actor
	Actor *ActorEntity `json:"actor,omitempty"`

	// response
	Response string `json:"response,omitempty"`
}

// Validate validates this pong entity
func (m *PongEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PongEntity) validateActor(formats strfmt.Registry) error {

	if swag.IsZero(m.Actor) { // not required
		return nil
	}

	if m.Actor != nil {
		if err := m.Actor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PongEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PongEntity) UnmarshalBinary(b []byte) error {
	var res PongEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
