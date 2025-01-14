// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
)

// ServiceDeleteHandlerFunc turns a function with the right signature into a service delete handler
type ServiceDeleteHandlerFunc func(ServiceDeleteParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceDeleteHandlerFunc) Handle(params ServiceDeleteParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ServiceDeleteHandler interface for that can handle valid service delete params
type ServiceDeleteHandler interface {
	Handle(ServiceDeleteParams, *models.Principal) middleware.Responder
}

// NewServiceDelete creates a new http.Handler for the service delete operation
func NewServiceDelete(ctx *middleware.Context, handler ServiceDeleteHandler) *ServiceDelete {
	return &ServiceDelete{Context: ctx, Handler: handler}
}

/*ServiceDelete swagger:route DELETE /services/{ServiceID}/ service serviceDelete

deletes a service item

Deletes a service object


*/
type ServiceDelete struct {
	Context *middleware.Context
	Handler ServiceDeleteHandler
}

func (o *ServiceDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewServiceDeleteParams()

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
