package postgresql

import (
	"fmt"
	apiv1 "k8s.io/api/core/v1"
)

const (
	postgresRoleKey       = "spilo-role"
	postgresRoleReplica   = "replica"
	postgresRoleMaster    = "master"
	postgresPodLabelKey   = "cluster-name"
	postgresPodDefaultKey = "application"
	postgresPodDefaultVal = "spilo"
	postgresMainContainer = "postgres"
	postgresPort          = 5432

	passwordField = "password"
)

type InternalDetails struct {
	Cluster *Postgres
}

func (d *InternalDetails) GetPodReplicaSelector() map[string]string {
	return map[string]string{postgresRoleKey: postgresRoleReplica,
		postgresPodLabelKey:   d.Cluster.Operator.ObjectMeta.Name,
		postgresPodDefaultKey: postgresPodDefaultVal,
	}
}

func (d *InternalDetails) GetPodMasterSelector() map[string]string {
	return map[string]string{postgresRoleKey: postgresRoleMaster,
		postgresPodLabelKey:   d.Cluster.Operator.ObjectMeta.Name,
		postgresPodDefaultKey: postgresPodDefaultVal,
	}
}

func (d *InternalDetails) GetMasterService() string {
	return fmt.Sprintf("%s", d.Cluster.Operator.ObjectMeta.Name)
}

func (d *InternalDetails) GetReplicaService() string {
	return fmt.Sprintf("%s-repl", d.Cluster.Operator.ObjectMeta.Name)
}

func (d *InternalDetails) GetAccessPort() int {
	return postgresPort
}

func (d *InternalDetails) GetMainPodContainer() string {
	return postgresMainContainer
}

func (d *InternalDetails) GetDefaultConnectionPassword() (string, string) {
	return genUserCredentialsSecretName(masterUser, d.Cluster.Operator.ObjectMeta.Name), passwordField
}

func (d *InternalDetails) GetDefaultConnectionUser() string {
	return masterUser
}

func (d *InternalDetails) GetCredentialsSecret() (*apiv1.Secret, error) {
	return nil, nil
}
