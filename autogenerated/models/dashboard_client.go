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

// DashboardClient dashboard client
// swagger:model DashboardClient
type DashboardClient struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// redirect uri
	RedirectURI string `json:"redirect_uri,omitempty"`

	// secret
	// Required: true
	Secret *string `json:"secret"`
}

// Validate validates this dashboard client
func (m *DashboardClient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardClient) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DashboardClient) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("secret", "body", m.Secret); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardClient) UnmarshalBinary(b []byte) error {
	var res DashboardClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
