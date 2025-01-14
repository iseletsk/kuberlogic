// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/models"
	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/restapi/operations/auth"
	"github.com/kuberlogic/kuberlogic/modules/apiserver/internal/generated/restapi/operations/service"
)

// NewKuberlogicAPI creates a new Kuberlogic instance
func NewKuberlogicAPI(spec *loads.Document) *KuberlogicAPI {
	return &KuberlogicAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ServiceBackupConfigCreateHandler: service.BackupConfigCreateHandlerFunc(func(params service.BackupConfigCreateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.BackupConfigCreate has not yet been implemented")
		}),
		ServiceBackupConfigDeleteHandler: service.BackupConfigDeleteHandlerFunc(func(params service.BackupConfigDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.BackupConfigDelete has not yet been implemented")
		}),
		ServiceBackupConfigEditHandler: service.BackupConfigEditHandlerFunc(func(params service.BackupConfigEditParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.BackupConfigEdit has not yet been implemented")
		}),
		ServiceBackupConfigGetHandler: service.BackupConfigGetHandlerFunc(func(params service.BackupConfigGetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.BackupConfigGet has not yet been implemented")
		}),
		ServiceBackupListHandler: service.BackupListHandlerFunc(func(params service.BackupListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.BackupList has not yet been implemented")
		}),
		ServiceDatabaseCreateHandler: service.DatabaseCreateHandlerFunc(func(params service.DatabaseCreateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.DatabaseCreate has not yet been implemented")
		}),
		ServiceDatabaseDeleteHandler: service.DatabaseDeleteHandlerFunc(func(params service.DatabaseDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.DatabaseDelete has not yet been implemented")
		}),
		ServiceDatabaseListHandler: service.DatabaseListHandlerFunc(func(params service.DatabaseListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.DatabaseList has not yet been implemented")
		}),
		ServiceDatabaseRestoreHandler: service.DatabaseRestoreHandlerFunc(func(params service.DatabaseRestoreParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.DatabaseRestore has not yet been implemented")
		}),
		AuthLoginUserHandler: auth.LoginUserHandlerFunc(func(params auth.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.LoginUser has not yet been implemented")
		}),
		ServiceLogsGetHandler: service.LogsGetHandlerFunc(func(params service.LogsGetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.LogsGet has not yet been implemented")
		}),
		ServiceRestoreListHandler: service.RestoreListHandlerFunc(func(params service.RestoreListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.RestoreList has not yet been implemented")
		}),
		ServiceServiceAddHandler: service.ServiceAddHandlerFunc(func(params service.ServiceAddParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceAdd has not yet been implemented")
		}),
		ServiceServiceDeleteHandler: service.ServiceDeleteHandlerFunc(func(params service.ServiceDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceDelete has not yet been implemented")
		}),
		ServiceServiceEditHandler: service.ServiceEditHandlerFunc(func(params service.ServiceEditParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceEdit has not yet been implemented")
		}),
		ServiceServiceGetHandler: service.ServiceGetHandlerFunc(func(params service.ServiceGetParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceGet has not yet been implemented")
		}),
		ServiceServiceListHandler: service.ServiceListHandlerFunc(func(params service.ServiceListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.ServiceList has not yet been implemented")
		}),
		ServiceUserCreateHandler: service.UserCreateHandlerFunc(func(params service.UserCreateParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.UserCreate has not yet been implemented")
		}),
		ServiceUserDeleteHandler: service.UserDeleteHandlerFunc(func(params service.UserDeleteParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.UserDelete has not yet been implemented")
		}),
		ServiceUserEditHandler: service.UserEditHandlerFunc(func(params service.UserEditParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.UserEdit has not yet been implemented")
		}),
		ServiceUserListHandler: service.UserListHandlerFunc(func(params service.UserListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation service.UserList has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*KuberlogicAPI This is a service API */
type KuberlogicAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (*models.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ServiceBackupConfigCreateHandler sets the operation handler for the backup config create operation
	ServiceBackupConfigCreateHandler service.BackupConfigCreateHandler
	// ServiceBackupConfigDeleteHandler sets the operation handler for the backup config delete operation
	ServiceBackupConfigDeleteHandler service.BackupConfigDeleteHandler
	// ServiceBackupConfigEditHandler sets the operation handler for the backup config edit operation
	ServiceBackupConfigEditHandler service.BackupConfigEditHandler
	// ServiceBackupConfigGetHandler sets the operation handler for the backup config get operation
	ServiceBackupConfigGetHandler service.BackupConfigGetHandler
	// ServiceBackupListHandler sets the operation handler for the backup list operation
	ServiceBackupListHandler service.BackupListHandler
	// ServiceDatabaseCreateHandler sets the operation handler for the database create operation
	ServiceDatabaseCreateHandler service.DatabaseCreateHandler
	// ServiceDatabaseDeleteHandler sets the operation handler for the database delete operation
	ServiceDatabaseDeleteHandler service.DatabaseDeleteHandler
	// ServiceDatabaseListHandler sets the operation handler for the database list operation
	ServiceDatabaseListHandler service.DatabaseListHandler
	// ServiceDatabaseRestoreHandler sets the operation handler for the database restore operation
	ServiceDatabaseRestoreHandler service.DatabaseRestoreHandler
	// AuthLoginUserHandler sets the operation handler for the login user operation
	AuthLoginUserHandler auth.LoginUserHandler
	// ServiceLogsGetHandler sets the operation handler for the logs get operation
	ServiceLogsGetHandler service.LogsGetHandler
	// ServiceRestoreListHandler sets the operation handler for the restore list operation
	ServiceRestoreListHandler service.RestoreListHandler
	// ServiceServiceAddHandler sets the operation handler for the service add operation
	ServiceServiceAddHandler service.ServiceAddHandler
	// ServiceServiceDeleteHandler sets the operation handler for the service delete operation
	ServiceServiceDeleteHandler service.ServiceDeleteHandler
	// ServiceServiceEditHandler sets the operation handler for the service edit operation
	ServiceServiceEditHandler service.ServiceEditHandler
	// ServiceServiceGetHandler sets the operation handler for the service get operation
	ServiceServiceGetHandler service.ServiceGetHandler
	// ServiceServiceListHandler sets the operation handler for the service list operation
	ServiceServiceListHandler service.ServiceListHandler
	// ServiceUserCreateHandler sets the operation handler for the user create operation
	ServiceUserCreateHandler service.UserCreateHandler
	// ServiceUserDeleteHandler sets the operation handler for the user delete operation
	ServiceUserDeleteHandler service.UserDeleteHandler
	// ServiceUserEditHandler sets the operation handler for the user edit operation
	ServiceUserEditHandler service.UserEditHandler
	// ServiceUserListHandler sets the operation handler for the user list operation
	ServiceUserListHandler service.UserListHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *KuberlogicAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *KuberlogicAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *KuberlogicAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *KuberlogicAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *KuberlogicAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *KuberlogicAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *KuberlogicAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *KuberlogicAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *KuberlogicAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the KuberlogicAPI
func (o *KuberlogicAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.ServiceBackupConfigCreateHandler == nil {
		unregistered = append(unregistered, "service.BackupConfigCreateHandler")
	}
	if o.ServiceBackupConfigDeleteHandler == nil {
		unregistered = append(unregistered, "service.BackupConfigDeleteHandler")
	}
	if o.ServiceBackupConfigEditHandler == nil {
		unregistered = append(unregistered, "service.BackupConfigEditHandler")
	}
	if o.ServiceBackupConfigGetHandler == nil {
		unregistered = append(unregistered, "service.BackupConfigGetHandler")
	}
	if o.ServiceBackupListHandler == nil {
		unregistered = append(unregistered, "service.BackupListHandler")
	}
	if o.ServiceDatabaseCreateHandler == nil {
		unregistered = append(unregistered, "service.DatabaseCreateHandler")
	}
	if o.ServiceDatabaseDeleteHandler == nil {
		unregistered = append(unregistered, "service.DatabaseDeleteHandler")
	}
	if o.ServiceDatabaseListHandler == nil {
		unregistered = append(unregistered, "service.DatabaseListHandler")
	}
	if o.ServiceDatabaseRestoreHandler == nil {
		unregistered = append(unregistered, "service.DatabaseRestoreHandler")
	}
	if o.AuthLoginUserHandler == nil {
		unregistered = append(unregistered, "auth.LoginUserHandler")
	}
	if o.ServiceLogsGetHandler == nil {
		unregistered = append(unregistered, "service.LogsGetHandler")
	}
	if o.ServiceRestoreListHandler == nil {
		unregistered = append(unregistered, "service.RestoreListHandler")
	}
	if o.ServiceServiceAddHandler == nil {
		unregistered = append(unregistered, "service.ServiceAddHandler")
	}
	if o.ServiceServiceDeleteHandler == nil {
		unregistered = append(unregistered, "service.ServiceDeleteHandler")
	}
	if o.ServiceServiceEditHandler == nil {
		unregistered = append(unregistered, "service.ServiceEditHandler")
	}
	if o.ServiceServiceGetHandler == nil {
		unregistered = append(unregistered, "service.ServiceGetHandler")
	}
	if o.ServiceServiceListHandler == nil {
		unregistered = append(unregistered, "service.ServiceListHandler")
	}
	if o.ServiceUserCreateHandler == nil {
		unregistered = append(unregistered, "service.UserCreateHandler")
	}
	if o.ServiceUserDeleteHandler == nil {
		unregistered = append(unregistered, "service.UserDeleteHandler")
	}
	if o.ServiceUserEditHandler == nil {
		unregistered = append(unregistered, "service.UserEditHandler")
	}
	if o.ServiceUserListHandler == nil {
		unregistered = append(unregistered, "service.UserListHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *KuberlogicAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *KuberlogicAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.BearerAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *KuberlogicAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *KuberlogicAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *KuberlogicAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *KuberlogicAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the kuberlogic API
func (o *KuberlogicAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *KuberlogicAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/services/{ServiceID}/backup-config"] = service.NewBackupConfigCreate(o.context, o.ServiceBackupConfigCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/services/{ServiceID}/backup-config"] = service.NewBackupConfigDelete(o.context, o.ServiceBackupConfigDeleteHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/services/{ServiceID}/backup-config"] = service.NewBackupConfigEdit(o.context, o.ServiceBackupConfigEditHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services/{ServiceID}/backup-config"] = service.NewBackupConfigGet(o.context, o.ServiceBackupConfigGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services/{ServiceID}/backups"] = service.NewBackupList(o.context, o.ServiceBackupListHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/services/{ServiceID}/databases"] = service.NewDatabaseCreate(o.context, o.ServiceDatabaseCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/services/{ServiceID}/databases/{Database}"] = service.NewDatabaseDelete(o.context, o.ServiceDatabaseDeleteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services/{ServiceID}/databases"] = service.NewDatabaseList(o.context, o.ServiceDatabaseListHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/services/{ServiceID}/restores"] = service.NewDatabaseRestore(o.context, o.ServiceDatabaseRestoreHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = auth.NewLoginUser(o.context, o.AuthLoginUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services/{ServiceID}/logs"] = service.NewLogsGet(o.context, o.ServiceLogsGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services/{ServiceID}/restores"] = service.NewRestoreList(o.context, o.ServiceRestoreListHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/services"] = service.NewServiceAdd(o.context, o.ServiceServiceAddHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/services/{ServiceID}"] = service.NewServiceDelete(o.context, o.ServiceServiceDeleteHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/services/{ServiceID}"] = service.NewServiceEdit(o.context, o.ServiceServiceEditHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services/{ServiceID}"] = service.NewServiceGet(o.context, o.ServiceServiceGetHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services"] = service.NewServiceList(o.context, o.ServiceServiceListHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/services/{ServiceID}/users"] = service.NewUserCreate(o.context, o.ServiceUserCreateHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/services/{ServiceID}/users/{Username}"] = service.NewUserDelete(o.context, o.ServiceUserDeleteHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/services/{ServiceID}/users/{Username}"] = service.NewUserEdit(o.context, o.ServiceUserEditHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/services/{ServiceID}/users"] = service.NewUserList(o.context, o.ServiceUserListHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *KuberlogicAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *KuberlogicAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *KuberlogicAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *KuberlogicAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *KuberlogicAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
