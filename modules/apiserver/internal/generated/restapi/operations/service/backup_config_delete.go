// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
)

// BackupConfigDeleteHandlerFunc turns a function with the right signature into a backup config delete handler
type BackupConfigDeleteHandlerFunc func(BackupConfigDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn BackupConfigDeleteHandlerFunc) Handle(params BackupConfigDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// BackupConfigDeleteHandler interface for that can handle valid backup config delete params
type BackupConfigDeleteHandler interface {
	Handle(BackupConfigDeleteParams, *models.Principal) middleware.Responder
}

// NewBackupConfigDelete creates a new http.Handler for the backup config delete operation
func NewBackupConfigDelete(ctx *middleware.Context, handler BackupConfigDeleteHandler) *BackupConfigDelete {
	return &BackupConfigDelete{Context: ctx, Handler: handler}
}

/*BackupConfigDelete swagger:route DELETE /services/{ServiceID}/backup-config service backupConfigDelete

BackupConfigDelete backup config delete API

*/
type BackupConfigDelete struct {
	Context *middleware.Context
	Handler BackupConfigDeleteHandler
}

func (o *BackupConfigDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBackupConfigDeleteParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}