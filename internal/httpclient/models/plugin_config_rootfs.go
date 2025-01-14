// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PluginConfigRootfs PluginConfigRootfs PluginConfigRootfs PluginConfigRootfs plugin config rootfs
//
// swagger:model PluginConfigRootfs
type PluginConfigRootfs struct {

	// diff ids
	DiffIds []string `json:"diff_ids"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this plugin config rootfs
func (m *PluginConfigRootfs) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this plugin config rootfs based on context it is used
func (m *PluginConfigRootfs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigRootfs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigRootfs) UnmarshalBinary(b []byte) error {
	var res PluginConfigRootfs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
