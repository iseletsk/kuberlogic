module github.com/kuberlogic/kuberlogic/modules/watcher

go 1.13

replace github.com/kuberlogic/kuberlogic/modules/operator => ../operator

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jackc/pgx/v4 v4.10.1
	github.com/kuberlogic/kuberlogic/modules/operator v0.0.21-0.20210519072214-3383b82a88f3
	github.com/pkg/errors v0.9.1
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v0.20.4
	sigs.k8s.io/controller-runtime v0.8.3
)
