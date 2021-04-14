// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImmutableSelector immutable selector
//
// swagger:model ImmutableSelector
type ImmutableSelector struct {

	// decoration
	Decoration string `json:"decoration,omitempty" js:"decoration"`

	// extras
	Extras string `json:"extras,omitempty" js:"extras"`

	// kind
	Kind string `json:"kind,omitempty" js:"kind"`

	// pattern
	Pattern string `json:"pattern,omitempty" js:"pattern"`
}

// Validate validates this immutable selector
func (m *ImmutableSelector) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this immutable selector based on context it is used
func (m *ImmutableSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImmutableSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImmutableSelector) UnmarshalBinary(b []byte) error {
	var res ImmutableSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}