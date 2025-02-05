// Code generated by go-swagger; DO NOT EDIT.
package service

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/security"
	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/logging/posthog"
	"github.com/kuberlogic/kuberlogic/modules/apiserver/util"
)

func UserCreateWrapper(srv Service, next UserCreateHandlerFunc) (fn UserCreateHandlerFunc) {
	return func(params UserCreateParams, principal *models.Principal) middleware.Responder {

		log := srv.GetLogger()

		// check ServiceID param
		ns, name, err := util.SplitID(params.ServiceID)
		if err != nil {
			msg := "incorrect service id"
			log.Errorw(msg, "serviceId", params.ServiceID, "error", err)
			return NewUserCreateBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		}

		// check auth
		authProvider := srv.GetAuthProvider()
		if authorized, err := authProvider.Authorize(principal, security.UserCreatePermission, params.ServiceID); err != nil {
			msg := "auth bad request"
			log.Errorw(msg, "permission", security.UserCreatePermission, "serviceId", params.ServiceID, "error", err)
			return NewUserCreateBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		} else if !authorized {
			log.Errorw("auth forbidden", "permission", security.UserCreatePermission, "serviceId", params.ServiceID)
			return NewUserCreateForbidden()
		}

		// cluster should exists
		service, found, err := srv.LookupService(ns, name)
		if !found {
			msg := "service not found"
			log.Errorw(msg, "error", err)
			return NewUserCreateBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		} else if err != nil {
			msg := "error getting service"
			log.Errorw(msg, "error", err)
			return NewUserCreateServiceUnavailable().WithPayload(&models.Error{
				Message: msg,
			})
		}

		params.HTTPRequest = params.HTTPRequest.WithContext(
			context.WithValue(params.HTTPRequest.Context(), "service", service))

		// enqueue data to posthog
		posthogMsg := posthog.NewMessage("user-create")
		posthogMsg.With("service-id", params.ServiceID)
		posthogMsg.With("user", params.User.Name)
		if perr := posthogMsg.Create(); perr != nil {
			msg := "could not enqueue posthog message"
			log.Errorw(msg, "error", perr)
			return NewUserCreateServiceUnavailable().WithPayload(&models.Error{
				Message: msg,
			})
		}

		return next(params, principal)
	}
}
