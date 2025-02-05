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

func ServiceEditWrapper(srv Service, next ServiceEditHandlerFunc) (fn ServiceEditHandlerFunc) {
	return func(params ServiceEditParams, principal *models.Principal) middleware.Responder {

		log := srv.GetLogger()

		// check ServiceID param
		ns, name, err := util.SplitID(params.ServiceID)
		if err != nil {
			msg := "incorrect service id"
			log.Errorw(msg, "serviceId", params.ServiceID, "error", err)
			return NewServiceEditBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		}

		// check auth
		authProvider := srv.GetAuthProvider()
		if authorized, err := authProvider.Authorize(principal, security.ServiceEditPermission, params.ServiceID); err != nil {
			msg := "auth bad request"
			log.Errorw(msg, "permission", security.ServiceEditPermission, "serviceId", params.ServiceID, "error", err)
			return NewServiceEditBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		} else if !authorized {
			log.Errorw("auth forbidden", "permission", security.ServiceEditPermission, "serviceId", params.ServiceID)
			return NewServiceEditForbidden()
		}

		// cluster should exists
		service, found, err := srv.LookupService(ns, name)
		if !found {
			msg := "service not found"
			log.Errorw(msg, "error", err)
			return NewServiceEditBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		} else if err != nil {
			msg := "error getting service"
			log.Errorw(msg, "error", err)
			return NewServiceEditServiceUnavailable().WithPayload(&models.Error{
				Message: msg,
			})
		}

		params.HTTPRequest = params.HTTPRequest.WithContext(
			context.WithValue(params.HTTPRequest.Context(), "service", service))

		// enqueue data to posthog
		posthogMsg := posthog.NewMessage("service-edit")
		posthogMsg.With("service-id", params.ServiceID)
		posthogMsg.With("type", params.ServiceItem.Type)

		if params.ServiceItem.Replicas != nil {
			posthogMsg.With("replicas", params.ServiceItem.Replicas)
		}

		if params.ServiceItem.Limits != nil && params.ServiceItem.Limits.CPU != nil {
			posthogMsg.With("cpu-limit", params.ServiceItem.Limits.CPU)
		}

		if params.ServiceItem.Limits != nil && params.ServiceItem.Limits.Memory != nil {
			posthogMsg.With("memory-limit", params.ServiceItem.Limits.Memory)
		}

		if params.ServiceItem.Limits != nil && params.ServiceItem.Limits.VolumeSize != nil {
			posthogMsg.With("volume-limit", params.ServiceItem.Limits.VolumeSize)
		}
		if perr := posthogMsg.Create(); perr != nil {
			msg := "could not enqueue posthog message"
			log.Errorw(msg, "error", perr)
			return NewServiceEditServiceUnavailable().WithPayload(&models.Error{
				Message: msg,
			})
		}

		return next(params, principal)
	}
}
