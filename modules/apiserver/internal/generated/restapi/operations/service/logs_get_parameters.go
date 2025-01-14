// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewLogsGetParams creates a new LogsGetParams object
// with the default values initialized.
func NewLogsGetParams() LogsGetParams {

	var (
		// initialize parameters with default values

		tailDefault = int64(100)
	)

	return LogsGetParams{
		Tail: &tailDefault,
	}
}

// LogsGetParams contains all the bound params for the logs get operation
// typically these are obtained from a http.Request
//
// swagger:parameters logsGet
type LogsGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*service Resource ID
	  Required: true
	  Max Length: 120
	  Min Length: 3
	  Pattern: [a-z0-9]([-a-z0-9]*[a-z0-9])?:[a-z0-9]([-a-z0-9]*[a-z0-9])?
	  In: path
	*/
	ServiceID string
	/*ServiceInstance to get logs of
	  Required: true
	  Pattern: ^\S+$
	  In: query
	*/
	ServiceInstance string
	/*number of lines to tail
	  In: query
	  Default: 100
	*/
	Tail *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewLogsGetParams() beforehand.
func (o *LogsGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rServiceID, rhkServiceID, _ := route.Params.GetOK("ServiceID")
	if err := o.bindServiceID(rServiceID, rhkServiceID, route.Formats); err != nil {
		res = append(res, err)
	}

	qServiceInstance, qhkServiceInstance, _ := qs.GetOK("service_instance")
	if err := o.bindServiceInstance(qServiceInstance, qhkServiceInstance, route.Formats); err != nil {
		res = append(res, err)
	}

	qTail, qhkTail, _ := qs.GetOK("tail")
	if err := o.bindTail(qTail, qhkTail, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindServiceID binds and validates parameter ServiceID from path.
func (o *LogsGetParams) bindServiceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ServiceID = raw

	if err := o.validateServiceID(formats); err != nil {
		return err
	}

	return nil
}

// validateServiceID carries on validations for parameter ServiceID
func (o *LogsGetParams) validateServiceID(formats strfmt.Registry) error {

	if err := validate.MinLength("ServiceID", "path", o.ServiceID, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("ServiceID", "path", o.ServiceID, 120); err != nil {
		return err
	}

	if err := validate.Pattern("ServiceID", "path", o.ServiceID, `[a-z0-9]([-a-z0-9]*[a-z0-9])?:[a-z0-9]([-a-z0-9]*[a-z0-9])?`); err != nil {
		return err
	}

	return nil
}

// bindServiceInstance binds and validates parameter ServiceInstance from query.
func (o *LogsGetParams) bindServiceInstance(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("service_instance", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("service_instance", "query", raw); err != nil {
		return err
	}

	o.ServiceInstance = raw

	if err := o.validateServiceInstance(formats); err != nil {
		return err
	}

	return nil
}

// validateServiceInstance carries on validations for parameter ServiceInstance
func (o *LogsGetParams) validateServiceInstance(formats strfmt.Registry) error {

	if err := validate.Pattern("service_instance", "query", o.ServiceInstance, `^\S+$`); err != nil {
		return err
	}

	return nil
}

// bindTail binds and validates parameter Tail from query.
func (o *LogsGetParams) bindTail(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewLogsGetParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("tail", "query", "int64", raw)
	}
	o.Tail = &value

	return nil
}
