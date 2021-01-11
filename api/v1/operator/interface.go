package operator

import (
	"github.com/pkg/errors"
	cloudlinuxv1 "gitlab.com/cloudmanaged/operator/api/v1"
	"gitlab.com/cloudmanaged/operator/api/v1/operator/mysql"
	"gitlab.com/cloudmanaged/operator/api/v1/operator/postgresql"
	"gitlab.com/cloudmanaged/operator/api/v1/operator/redis"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Operator interface {
	Init(cm *cloudlinuxv1.CloudManaged)
	InitFrom(o runtime.Object)
	Update(cm *cloudlinuxv1.CloudManaged)
	AsRuntimeObject() runtime.Object
	AsMetaObject() metav1.Object
	IsEqual(cm *cloudlinuxv1.CloudManaged) bool
	CurrentStatus() string
	GetDefaults() cloudlinuxv1.Defaults

	GetPodReplicaSelector(string) map[string]string
	GetPodMasterSelector(string) map[string]string

	GetMasterService(string, string) string
	GetReplicaService(string, string) string
}

type Backup interface {
	New(backup *cloudlinuxv1.CloudManagedBackup) v1beta1.CronJob
	Init(*cloudlinuxv1.CloudManagedBackup)
	InitFrom(*v1beta1.CronJob)
	IsEqual(cm *cloudlinuxv1.CloudManagedBackup) bool
	Update(cm *cloudlinuxv1.CloudManagedBackup)
	GetCronJob() *v1beta1.CronJob
	CurrentStatus(ev batchv1.JobList) string

	SetBackupImage()
	SetBackupEnv(cm *cloudlinuxv1.CloudManagedBackup)
}

type Restore interface {
	New(backup *cloudlinuxv1.CloudManagedRestore) batchv1.Job
	Init(*cloudlinuxv1.CloudManagedRestore)
	InitFrom(*batchv1.Job)
	GetJob() *batchv1.Job
	CurrentStatus() string

	SetRestoreImage()
	SetRestoreEnv(cm *cloudlinuxv1.CloudManagedRestore)
}

func GetOperator(t string) (Operator, error) {
	var operators = map[string]Operator{
		"postgresql": &postgresql.Postgres{},
		"redis":      &redis.Redis{},
		"mysql":      &mysql.Mysql{},
	}

	value, ok := operators[t]
	if !ok {
		return nil, errors.Errorf("Operator %s is not supported", t)
	}
	return value, nil
}

func GetBackupOperator(op interface{}) (Backup, error) {
	switch cluster := op.(type) {
	case *mysql.Mysql:
		return &mysql.Backup{
			Cluster: *cluster,
		}, nil
	case *postgresql.Postgres:
		return &postgresql.Backup{
			Cluster: *cluster,
		}, nil
	default:
		return nil, errors.Errorf("Cluster %s is not supported (%T)",
			cluster, cluster)
	}
}

func GetRestoreOperator(op interface{}) (Restore, error) {
	switch cluster := op.(type) {
	case *mysql.Mysql:
		return &mysql.Restore{
			Cluster: *cluster,
		}, nil
	case *postgresql.Postgres:
		return &postgresql.Restore{
			Cluster: *cluster,
		}, nil
	default:
		return nil, errors.Errorf("Cluster %s is not supported (%T)",
			cluster, cluster)
	}
}
