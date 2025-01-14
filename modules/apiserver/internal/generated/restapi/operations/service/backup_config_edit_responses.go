// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
)

// BackupConfigEditOKCode is the HTTP code returned for type BackupConfigEditOK
const BackupConfigEditOKCode int = 200

/*BackupConfigEditOK item is edited

swagger:response backupConfigEditOK
*/
type BackupConfigEditOK struct {

	/*
	  In: Body
	*/
	Payload *models.BackupConfig `json:"body,omitempty"`
}

// NewBackupConfigEditOK creates BackupConfigEditOK with default headers values
func NewBackupConfigEditOK() *BackupConfigEditOK {

	return &BackupConfigEditOK{}
}

// WithPayload adds the payload to the backup config edit o k response
func (o *BackupConfigEditOK) WithPayload(payload *models.BackupConfig) *BackupConfigEditOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backup config edit o k response
func (o *BackupConfigEditOK) SetPayload(payload *models.BackupConfig) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupConfigEditOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BackupConfigEditBadRequestCode is the HTTP code returned for type BackupConfigEditBadRequest
const BackupConfigEditBadRequestCode int = 400

/*BackupConfigEditBadRequest invalid input, object invalid

swagger:response backupConfigEditBadRequest
*/
type BackupConfigEditBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBackupConfigEditBadRequest creates BackupConfigEditBadRequest with default headers values
func NewBackupConfigEditBadRequest() *BackupConfigEditBadRequest {

	return &BackupConfigEditBadRequest{}
}

// WithPayload adds the payload to the backup config edit bad request response
func (o *BackupConfigEditBadRequest) WithPayload(payload *models.Error) *BackupConfigEditBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backup config edit bad request response
func (o *BackupConfigEditBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupConfigEditBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BackupConfigEditUnauthorizedCode is the HTTP code returned for type BackupConfigEditUnauthorized
const BackupConfigEditUnauthorizedCode int = 401

/*BackupConfigEditUnauthorized bad authentication

swagger:response backupConfigEditUnauthorized
*/
type BackupConfigEditUnauthorized struct {
}

// NewBackupConfigEditUnauthorized creates BackupConfigEditUnauthorized with default headers values
func NewBackupConfigEditUnauthorized() *BackupConfigEditUnauthorized {

	return &BackupConfigEditUnauthorized{}
}

// WriteResponse to the client
func (o *BackupConfigEditUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// BackupConfigEditForbiddenCode is the HTTP code returned for type BackupConfigEditForbidden
const BackupConfigEditForbiddenCode int = 403

/*BackupConfigEditForbidden bad permissions

swagger:response backupConfigEditForbidden
*/
type BackupConfigEditForbidden struct {
}

// NewBackupConfigEditForbidden creates BackupConfigEditForbidden with default headers values
func NewBackupConfigEditForbidden() *BackupConfigEditForbidden {

	return &BackupConfigEditForbidden{}
}

// WriteResponse to the client
func (o *BackupConfigEditForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// BackupConfigEditNotFoundCode is the HTTP code returned for type BackupConfigEditNotFound
const BackupConfigEditNotFoundCode int = 404

/*BackupConfigEditNotFound item not found

swagger:response backupConfigEditNotFound
*/
type BackupConfigEditNotFound struct {
}

// NewBackupConfigEditNotFound creates BackupConfigEditNotFound with default headers values
func NewBackupConfigEditNotFound() *BackupConfigEditNotFound {

	return &BackupConfigEditNotFound{}
}

// WriteResponse to the client
func (o *BackupConfigEditNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// BackupConfigEditServiceUnavailableCode is the HTTP code returned for type BackupConfigEditServiceUnavailable
const BackupConfigEditServiceUnavailableCode int = 503

/*BackupConfigEditServiceUnavailable internal server error

swagger:response backupConfigEditServiceUnavailable
*/
type BackupConfigEditServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBackupConfigEditServiceUnavailable creates BackupConfigEditServiceUnavailable with default headers values
func NewBackupConfigEditServiceUnavailable() *BackupConfigEditServiceUnavailable {

	return &BackupConfigEditServiceUnavailable{}
}

// WithPayload adds the payload to the backup config edit service unavailable response
func (o *BackupConfigEditServiceUnavailable) WithPayload(payload *models.Error) *BackupConfigEditServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the backup config edit service unavailable response
func (o *BackupConfigEditServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BackupConfigEditServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
