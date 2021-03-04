// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
)

// BackupListHandlerFunc turns a function with the right signature into a backup list handler
type BackupListHandlerFunc func(BackupListParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn BackupListHandlerFunc) Handle(params BackupListParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// BackupListHandler interface for that can handle valid backup list params
type BackupListHandler interface {
	Handle(BackupListParams, *models.Principal) middleware.Responder
}

// NewBackupList creates a new http.Handler for the backup list operation
func NewBackupList(ctx *middleware.Context, handler BackupListHandler) *BackupList {
	return &BackupList{Context: ctx, Handler: handler}
}

/*BackupList swagger:route GET /services/{ServiceID}/backups/ service backupList

BackupList backup list API

*/
type BackupList struct {
	Context *middleware.Context
	Handler BackupListHandler
}

func (o *BackupList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBackupListParams()

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
