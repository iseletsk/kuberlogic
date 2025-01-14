// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
)

// RestoreListHandlerFunc turns a function with the right signature into a restore list handler
type RestoreListHandlerFunc func(RestoreListParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn RestoreListHandlerFunc) Handle(params RestoreListParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// RestoreListHandler interface for that can handle valid restore list params
type RestoreListHandler interface {
	Handle(RestoreListParams, *models.Principal) middleware.Responder
}

// NewRestoreList creates a new http.Handler for the restore list operation
func NewRestoreList(ctx *middleware.Context, handler RestoreListHandler) *RestoreList {
	return &RestoreList{Context: ctx, Handler: handler}
}

/*RestoreList swagger:route GET /services/{ServiceID}/restores service restoreList

RestoreList restore list API

*/
type RestoreList struct {
	Context *middleware.Context
	Handler RestoreListHandler
}

func (o *RestoreList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRestoreListParams()

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
