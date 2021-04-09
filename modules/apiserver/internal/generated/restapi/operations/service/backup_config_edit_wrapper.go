// Code generated by go-swagger; DO NOT EDIT.
package service

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/security"
	"github.com/kuberlogic/operator/modules/apiserver/internal/logging/posthog"
	"github.com/kuberlogic/operator/modules/apiserver/util"
)

func BackupConfigEditWrapper(srv Service, next BackupConfigEditHandlerFunc) (fn BackupConfigEditHandlerFunc) {
	return func(params BackupConfigEditParams, principal *models.Principal) middleware.Responder {

		log := srv.GetLogger()

		// check ServiceID param
		ns, name, err := util.SplitID(params.ServiceID)
		if err != nil {
			msg := "incorrect service id"
			log.Errorw(msg, "serviceId", params.ServiceID, "error", err)
			return NewBackupConfigEditBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		}

		// check auth
		authProvider := srv.GetAuthProvider()
		if authorized, err := authProvider.Authorize(principal.Token, security.BackupConfigEditPermission, params.ServiceID); err != nil {
			msg := "auth bad request"
			log.Errorw(msg, "permission", security.BackupConfigEditPermission, "serviceId", params.ServiceID, "error", err)
			return NewBackupConfigEditBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		} else if !authorized {
			log.Errorw("auth forbidden", "permission", security.BackupConfigEditPermission, "serviceId", params.ServiceID)
			return NewBackupConfigEditForbidden()
		}

		// cluster should exists
		err = srv.LookupService(ns, name)
		if err != nil {
			msg := "service does not exist"
			log.Errorw(msg, "error", err)
			return NewBackupConfigEditBadRequest().WithPayload(&models.Error{
				Message: msg,
			})
		}

		// enqueue data to posthog
		posthogMsg := posthog.NewMessage("backup-config-edit")
		posthogMsg.With("service-id", params.ServiceID)
		posthogMsg.With("enabled", params.BackupConfig.Enabled)
		posthogMsg.With("schedule", params.BackupConfig.Schedule)
		posthogMsg.With("bucket", params.BackupConfig.Bucket)
		posthogMsg.With("endpoint", params.BackupConfig.Endpoint)
		if perr := posthogMsg.Create(); perr != nil {
			msg := "could not enqueue posthog message"
			log.Errorw(msg, "error", perr)
			return NewBackupConfigEditServiceUnavailable().WithPayload(&models.Error{
				Message: msg,
			})
		}

		return next(params, principal)
	}
}
