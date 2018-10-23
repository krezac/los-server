// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Range range
// swagger:model Range
type Range struct {

	// shooting range active
	Active bool `json:"active,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// latitude
	Latitude float64 `json:"latitude,omitempty"`

	// longitude
	Longitude float64 `json:"longitude,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this range
func (m *Range) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Range) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Range) UnmarshalBinary(b []byte) error {
	var res Range
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
