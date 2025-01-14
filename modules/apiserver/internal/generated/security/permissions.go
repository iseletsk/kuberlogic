// Code generated by go-swagger; DO NOT EDIT.
package security

const (
	BackupConfigCreatePermission = "service:backup-config:add"
	BackupConfigDeletePermission = "service:backup-config:delete"
	BackupConfigEditPermission   = "service:backup-config:edit"
	BackupConfigGetPermission    = "service:backup-config:get"
	BackupListPermission         = "service:backup:list"
	DatabaseCreatePermission     = "service:database:add"
	DatabaseDeletePermission     = "service:database:delete"
	DatabaseListPermission       = "service:database:list"
	DatabaseRestorePermission    = "service:backup-restore:add"

	LogsGetPermission       = "service:logs"
	RestoreListPermission   = "service:restore:list"
	ServiceAddPermission    = "service:add"
	ServiceDeletePermission = "service:delete"
	ServiceEditPermission   = "service:edit"
	ServiceGetPermission    = "service:get"
	ServiceListPermission   = "service:list"
	UserCreatePermission    = "service:user:add"
	UserDeletePermission    = "service:user:delete"
	UserEditPermission      = "service:user:edit"
	UserListPermission      = "service:user:list"
)

var (
	ServiceGrants = []string{
		BackupConfigCreatePermission,
		BackupConfigDeletePermission,
		BackupConfigEditPermission,
		BackupConfigGetPermission,
		BackupListPermission,
		DatabaseCreatePermission,
		DatabaseDeletePermission,
		DatabaseListPermission,
		DatabaseRestorePermission,

		LogsGetPermission,
		RestoreListPermission,
		ServiceAddPermission,
		ServiceDeletePermission,
		ServiceEditPermission,
		ServiceGetPermission,
		ServiceListPermission,
		UserCreatePermission,
		UserDeletePermission,
		UserEditPermission,
		UserListPermission,
	}
)
