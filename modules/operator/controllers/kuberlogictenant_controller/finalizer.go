package kuberlogictenant_controller

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	kuberlogicv1 "github.com/kuberlogic/operator/modules/operator/api/v1"
	"github.com/kuberlogic/operator/modules/operator/cfg"
	"github.com/kuberlogic/operator/modules/operator/controllers/kuberlogictenant_controller/grafana"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// finalize function "resolves" an alert when kuberlogictenant is deleted.
func finalize(ctx context.Context, cfg *cfg.Config, c client.Client, kt *kuberlogicv1.KuberLogicTenant, log logr.Logger) error {
	log.Info("processing finalizer")
	if cfg.Grafana.Enabled {
		log.Info("processing grafana organizations and users")
		if err := grafana.NewGrafanaSyncer(kt, log, cfg.Grafana).DeleteOrganizationAndUsers(kt.Name); err != nil {
			return err
		}
	}

	log.Info("processing kuberlogic services")
	klsList := new(kuberlogicv1.KuberLogicServiceList)
	ns := kt.GetTenantName()

	listOption := &client.ListOptions{
		Namespace: ns,
	}
	if err := c.List(ctx, klsList, listOption); err != nil {
		log.Error(err, "error listing kuberlogicservices for tenant", "namespace", kt.GetNamespace())
	}

	log.Info("checking if a tenant namespace is empty", "kuberlogicservices", len(klsList.Items))
	if len(klsList.Items) != 0 {
		return fmt.Errorf("tenant must be cleaned up before deletion")
	}
	return nil
}
