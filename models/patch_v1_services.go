// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchV1Services Update a services attributes, you may also add or remove functionalities from the service as well.
// Note: You may not remove or add individual label key/value pairs. You must include the entire object to override label values.
//
// swagger:model patchV1Services
type PatchV1Services struct {

	// description
	Description string `json:"description,omitempty"`

	// An array of functionalities
	Functionalities []*PatchV1ServicesFunctionalitiesItems0 `json:"functionalities"`

	// A hash of label keys and values
	Labels map[string]string `json:"labels,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// If set to true, any functionality objects tagged on the service that are not included in the given array will be removed. Set this if you want to do a replacement operation for the functionalities
	RemoveRemainingFunctionalities bool `json:"remove_remaining_functionalities,omitempty"`
}

// Validate validates this patch v1 services
func (m *PatchV1Services) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunctionalities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1Services) validateFunctionalities(formats strfmt.Registry) error {

	if swag.IsZero(m.Functionalities) { // not required
		return nil
	}

	for i := 0; i < len(m.Functionalities); i++ {
		if swag.IsZero(m.Functionalities[i]) { // not required
			continue
		}

		if m.Functionalities[i] != nil {
			if err := m.Functionalities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("functionalities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchV1Services) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1Services) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1Services) UnmarshalBinary(b []byte) error {
	var res PatchV1Services
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1ServicesFunctionalitiesItems0 patch v1 services functionalities items0
// swagger:model PatchV1ServicesFunctionalitiesItems0
type PatchV1ServicesFunctionalitiesItems0 struct {

	// If you are trying to reuse a functionality, you may set the ID to attach it to the service
	ID string `json:"id,omitempty"`

	// If you are trying to remove a functionality from a service, set this to "true"
	Remove bool `json:"remove,omitempty"`

	// If you are trying to create a new functionality and attach it to this service, set the summary key
	Summary string `json:"summary,omitempty"`
}

// Validate validates this patch v1 services functionalities items0
func (m *PatchV1ServicesFunctionalitiesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1ServicesFunctionalitiesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1ServicesFunctionalitiesItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1ServicesFunctionalitiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
