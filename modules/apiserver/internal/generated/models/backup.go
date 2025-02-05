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

// Backup backup
//
// swagger:model Backup
type Backup struct {

	// file
	// Required: true
	File *string `json:"file"`

	// last modified
	// Required: true
	// Format: date-time
	LastModified *strfmt.DateTime `json:"lastModified"`

	// size
	// Required: true
	Size *int64 `json:"size"`
}

// Validate validates this backup
func (m *Backup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Backup) validateFile(formats strfmt.Registry) error {

	if err := validate.Required("file", "body", m.File); err != nil {
		return err
	}

	return nil
}

func (m *Backup) validateLastModified(formats strfmt.Registry) error {

	if err := validate.Required("lastModified", "body", m.LastModified); err != nil {
		return err
	}

	if err := validate.FormatOf("lastModified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Backup) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Backup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Backup) UnmarshalBinary(b []byte) error {
	var res Backup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
