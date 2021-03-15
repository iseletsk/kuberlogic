// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Limits limits
//
// swagger:model Limits
type Limits struct {

	// cpu
	// Required: true
	CPU *string `json:"cpu"`

	// memory
	// Required: true
	// Pattern: ^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$
	Memory *string `json:"memory"`

	// volume size
	// Required: true
	// Pattern: ^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$
	VolumeSize *string `json:"volumeSize"`
}

// Validate validates this limits
func (m *Limits) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Limits) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("cpu", "body", m.CPU); err != nil {
		return err
	}

	return nil
}

func (m *Limits) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	if err := validate.Pattern("memory", "body", string(*m.Memory), `^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`); err != nil {
		return err
	}

	return nil
}

func (m *Limits) validateVolumeSize(formats strfmt.Registry) error {

	if err := validate.Required("volumeSize", "body", m.VolumeSize); err != nil {
		return err
	}

	if err := validate.Pattern("volumeSize", "body", string(*m.VolumeSize), `^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Limits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Limits) UnmarshalBinary(b []byte) error {
	var res Limits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}