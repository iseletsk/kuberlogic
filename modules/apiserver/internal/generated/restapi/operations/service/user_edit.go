// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
)

// UserEditHandlerFunc turns a function with the right signature into a user edit handler
type UserEditHandlerFunc func(UserEditParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UserEditHandlerFunc) Handle(params UserEditParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UserEditHandler interface for that can handle valid user edit params
type UserEditHandler interface {
	Handle(UserEditParams, *models.Principal) middleware.Responder
}

// NewUserEdit creates a new http.Handler for the user edit operation
func NewUserEdit(ctx *middleware.Context, handler UserEditHandler) *UserEdit {
	return &UserEdit{Context: ctx, Handler: handler}
}

/*UserEdit swagger:route PUT /services/{ServiceID}/users/{Username}/ service userEdit

UserEdit user edit API

*/
type UserEdit struct {
	Context *middleware.Context
	Handler UserEditHandler
}

func (o *UserEdit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserEditParams()

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
