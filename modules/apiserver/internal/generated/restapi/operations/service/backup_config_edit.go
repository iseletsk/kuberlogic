// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
)

// BackupConfigEditHandlerFunc turns a function with the right signature into a backup config edit handler
type BackupConfigEditHandlerFunc func(BackupConfigEditParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn BackupConfigEditHandlerFunc) Handle(params BackupConfigEditParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// BackupConfigEditHandler interface for that can handle valid backup config edit params
type BackupConfigEditHandler interface {
	Handle(BackupConfigEditParams, *models.Principal) middleware.Responder
}

// NewBackupConfigEdit creates a new http.Handler for the backup config edit operation
func NewBackupConfigEdit(ctx *middleware.Context, handler BackupConfigEditHandler) *BackupConfigEdit {
	return &BackupConfigEdit{Context: ctx, Handler: handler}
}

/*BackupConfigEdit swagger:route PUT /services/{ServiceID}/backup-config service backupConfigEdit

BackupConfigEdit backup config edit API

*/
type BackupConfigEdit struct {
	Context *middleware.Context
	Handler BackupConfigEditHandler
}

func (o *BackupConfigEdit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBackupConfigEditParams()

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
