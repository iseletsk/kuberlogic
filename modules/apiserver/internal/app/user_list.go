package app

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/kuberlogic/operator/modules/apiserver/internal/generated/models"
	apiService "github.com/kuberlogic/operator/modules/apiserver/internal/generated/restapi/operations/service"
	"github.com/kuberlogic/operator/modules/apiserver/util"
	kuberlogicv1 "github.com/kuberlogic/operator/modules/operator/api/v1"
	"github.com/kuberlogic/operator/modules/operator/service-operator/util/kuberlogic"
)

// set this string to a required security grant for this action
const userListSecGrant = "service:user:list"

func (srv *Service) UserListHandler(params apiService.UserListParams, principal *models.Principal) middleware.Responder {
	// validate path parameter
	ns, name, err := util.SplitID(params.ServiceID)
	if err != nil {
		srv.log.Errorf("incorrect service id: %s", err.Error())
		return util.BadRequestFromError(err)
	}

	if authorized, err := srv.authProvider.Authorize(principal.Token, userListSecGrant, params.ServiceID); err != nil {
		srv.log.Errorf("error checking authorization: %s", err.Error())
		resp := apiService.NewDatabaseListBadRequest()
		return resp
	} else if !authorized {
		resp := apiService.NewDatabaseListForbidden()
		return resp
	}

	// check cluster is exists
	item := kuberlogicv1.KuberLogicService{}
	err = srv.cmClient.Get().
		Namespace(ns).
		Resource("kuberlogicservices").
		Name(name).
		Do(context.TODO()).
		Into(&item)
	if err != nil {
		srv.log.Errorf("couldn't find KuberLogicService resource in cluster: %s", err.Error())
		return util.BadRequestFromError(err)
	}

	session, err := kuberlogic.GetSession(&item, srv.clientset, "")
	if err != nil {
		srv.log.Errorf("error generating session: %s", err.Error())
		return util.BadRequestFromError(err)
	}

	users, err := session.GetUser().List()
	if err != nil {
		srv.log.Errorf("error receiving databases: %s", err.Error())
		return util.BadRequestFromError(err)
	}

	var payload []*models.User
	for _, dbUser := range users {
		userName := dbUser

		if protected := session.GetUser().IsProtected(userName); !protected {
			payload = append(payload, &models.User{
				Name: &userName,
			})
		}
	}

	return apiService.NewUserListOK().WithPayload(payload)
}
