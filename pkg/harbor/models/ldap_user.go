// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LdapUser ldap user
//
// swagger:model LdapUser
type LdapUser struct {

	// The user email address from "mail" or "email" attribute.
	Email string `json:"email,omitempty" js:"email"`

	// The user realname from "uid" or "cn" attribute.
	Realname string `json:"realname,omitempty" js:"realname"`

	// ldap username.
	Username string `json:"username,omitempty" js:"username"`
}

// Validate validates this ldap user
func (m *LdapUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ldap user based on context it is used
func (m *LdapUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapUser) UnmarshalBinary(b []byte) error {
	var res LdapUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}