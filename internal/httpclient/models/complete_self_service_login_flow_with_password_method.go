// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod CompleteSelfServiceLoginFlowWithPasswordMethod complete self service login flow with password method
//
// swagger:model CompleteSelfServiceLoginFlowWithPasswordMethod
type CompleteSelfServiceLoginFlowWithPasswordMethod struct {

	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken string `json:"csrf_token,omitempty"`

	// Identifier is the email or username of the user trying to log in.
	Identifier string `json:"identifier,omitempty"`

	// The user's password.
	Password string `json:"password,omitempty"`
}

// Validate validates this complete self service login flow with password method
func (m *CompleteSelfServiceLoginFlowWithPasswordMethod) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this complete self service login flow with password method based on context it is used
func (m *CompleteSelfServiceLoginFlowWithPasswordMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompleteSelfServiceLoginFlowWithPasswordMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompleteSelfServiceLoginFlowWithPasswordMethod) UnmarshalBinary(b []byte) error {
	var res CompleteSelfServiceLoginFlowWithPasswordMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
