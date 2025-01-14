// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
)

// UserListHandlerFunc turns a function with the right signature into a user list handler
type UserListHandlerFunc func(UserListParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UserListHandlerFunc) Handle(params UserListParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UserListHandler interface for that can handle valid user list params
type UserListHandler interface {
	Handle(UserListParams, *models.Principal) middleware.Responder
}

// NewUserList creates a new http.Handler for the user list operation
func NewUserList(ctx *middleware.Context, handler UserListHandler) *UserList {
	return &UserList{Context: ctx, Handler: handler}
}

/*UserList swagger:route GET /services/{ServiceID}/users/ service userList

UserList user list API

*/
type UserList struct {
	Context *middleware.Context
	Handler UserListHandler
}

func (o *UserList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserListParams()

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
