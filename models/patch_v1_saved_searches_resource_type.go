// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PatchV1SavedSearchesResourceType Update a specific saved search
// swagger:model patchV1SavedSearchesResourceType
type PatchV1SavedSearchesResourceType struct {

	// filter values
	FilterValues interface{} `json:"filter_values,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this patch v1 saved searches resource type
func (m *PatchV1SavedSearchesResourceType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1SavedSearchesResourceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1SavedSearchesResourceType) UnmarshalBinary(b []byte) error {
	var res PatchV1SavedSearchesResourceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
