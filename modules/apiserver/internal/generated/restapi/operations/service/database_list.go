// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
)

// DatabaseListHandlerFunc turns a function with the right signature into a database list handler
type DatabaseListHandlerFunc func(DatabaseListParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DatabaseListHandlerFunc) Handle(params DatabaseListParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DatabaseListHandler interface for that can handle valid database list params
type DatabaseListHandler interface {
	Handle(DatabaseListParams, *models.Principal) middleware.Responder
}

// NewDatabaseList creates a new http.Handler for the database list operation
func NewDatabaseList(ctx *middleware.Context, handler DatabaseListHandler) *DatabaseList {
	return &DatabaseList{Context: ctx, Handler: handler}
}

/*DatabaseList swagger:route GET /services/{ServiceID}/databases/ service databaseList

DatabaseList database list API

*/
type DatabaseList struct {
	Context *middleware.Context
	Handler DatabaseListHandler
}

func (o *DatabaseList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDatabaseListParams()

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
