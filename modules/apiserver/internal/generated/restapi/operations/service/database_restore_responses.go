// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
)

// DatabaseRestoreOKCode is the HTTP code returned for type DatabaseRestoreOK
const DatabaseRestoreOKCode int = 200

/*DatabaseRestoreOK item is restored

swagger:response databaseRestoreOK
*/
type DatabaseRestoreOK struct {
}

// NewDatabaseRestoreOK creates DatabaseRestoreOK with default headers values
func NewDatabaseRestoreOK() *DatabaseRestoreOK {

	return &DatabaseRestoreOK{}
}

// WriteResponse to the client
func (o *DatabaseRestoreOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DatabaseRestoreBadRequestCode is the HTTP code returned for type DatabaseRestoreBadRequest
const DatabaseRestoreBadRequestCode int = 400

/*DatabaseRestoreBadRequest invalid input, object invalid

swagger:response databaseRestoreBadRequest
*/
type DatabaseRestoreBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDatabaseRestoreBadRequest creates DatabaseRestoreBadRequest with default headers values
func NewDatabaseRestoreBadRequest() *DatabaseRestoreBadRequest {

	return &DatabaseRestoreBadRequest{}
}

// WithPayload adds the payload to the database restore bad request response
func (o *DatabaseRestoreBadRequest) WithPayload(payload *models.Error) *DatabaseRestoreBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the database restore bad request response
func (o *DatabaseRestoreBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DatabaseRestoreBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DatabaseRestoreUnauthorizedCode is the HTTP code returned for type DatabaseRestoreUnauthorized
const DatabaseRestoreUnauthorizedCode int = 401

/*DatabaseRestoreUnauthorized bad authentication

swagger:response databaseRestoreUnauthorized
*/
type DatabaseRestoreUnauthorized struct {
}

// NewDatabaseRestoreUnauthorized creates DatabaseRestoreUnauthorized with default headers values
func NewDatabaseRestoreUnauthorized() *DatabaseRestoreUnauthorized {

	return &DatabaseRestoreUnauthorized{}
}

// WriteResponse to the client
func (o *DatabaseRestoreUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DatabaseRestoreForbiddenCode is the HTTP code returned for type DatabaseRestoreForbidden
const DatabaseRestoreForbiddenCode int = 403

/*DatabaseRestoreForbidden bad permissions

swagger:response databaseRestoreForbidden
*/
type DatabaseRestoreForbidden struct {
}

// NewDatabaseRestoreForbidden creates DatabaseRestoreForbidden with default headers values
func NewDatabaseRestoreForbidden() *DatabaseRestoreForbidden {

	return &DatabaseRestoreForbidden{}
}

// WriteResponse to the client
func (o *DatabaseRestoreForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DatabaseRestoreNotFoundCode is the HTTP code returned for type DatabaseRestoreNotFound
const DatabaseRestoreNotFoundCode int = 404

/*DatabaseRestoreNotFound item not found

swagger:response databaseRestoreNotFound
*/
type DatabaseRestoreNotFound struct {
}

// NewDatabaseRestoreNotFound creates DatabaseRestoreNotFound with default headers values
func NewDatabaseRestoreNotFound() *DatabaseRestoreNotFound {

	return &DatabaseRestoreNotFound{}
}

// WriteResponse to the client
func (o *DatabaseRestoreNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DatabaseRestoreServiceUnavailableCode is the HTTP code returned for type DatabaseRestoreServiceUnavailable
const DatabaseRestoreServiceUnavailableCode int = 503

/*DatabaseRestoreServiceUnavailable internal server error

swagger:response databaseRestoreServiceUnavailable
*/
type DatabaseRestoreServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDatabaseRestoreServiceUnavailable creates DatabaseRestoreServiceUnavailable with default headers values
func NewDatabaseRestoreServiceUnavailable() *DatabaseRestoreServiceUnavailable {

	return &DatabaseRestoreServiceUnavailable{}
}

// WithPayload adds the payload to the database restore service unavailable response
func (o *DatabaseRestoreServiceUnavailable) WithPayload(payload *models.Error) *DatabaseRestoreServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the database restore service unavailable response
func (o *DatabaseRestoreServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DatabaseRestoreServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
