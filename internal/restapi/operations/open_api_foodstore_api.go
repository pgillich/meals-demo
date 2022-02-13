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

	"github.com/pgillich/meals-demo/internal/models"
	"github.com/pgillich/meals-demo/internal/restapi/operations/info"
	"github.com/pgillich/meals-demo/internal/restapi/operations/meal"
	"github.com/pgillich/meals-demo/internal/restapi/operations/user"
)

// NewOpenAPIFoodstoreAPI creates a new OpenAPIFoodstore instance
func NewOpenAPIFoodstoreAPI(spec *loads.Document) *OpenAPIFoodstoreAPI {
	return &OpenAPIFoodstoreAPI{
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

		MealCreateMealHandler: meal.CreateMealHandlerFunc(func(params meal.CreateMealParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation meal.CreateMeal has not yet been implemented")
		}),
		MealDeleteMealHandler: meal.DeleteMealHandlerFunc(func(params meal.DeleteMealParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation meal.DeleteMeal has not yet been implemented")
		}),
		MealFindMealsByTagHandler: meal.FindMealsByTagHandlerFunc(func(params meal.FindMealsByTagParams) middleware.Responder {
			return middleware.NotImplemented("operation meal.FindMealsByTag has not yet been implemented")
		}),
		MealGetIngredientsHandler: meal.GetIngredientsHandlerFunc(func(params meal.GetIngredientsParams) middleware.Responder {
			return middleware.NotImplemented("operation meal.GetIngredients has not yet been implemented")
		}),
		InfoGetLivezHandler: info.GetLivezHandlerFunc(func(params info.GetLivezParams) middleware.Responder {
			return middleware.NotImplemented("operation info.GetLivez has not yet been implemented")
		}),
		MealGetMealByIDHandler: meal.GetMealByIDHandlerFunc(func(params meal.GetMealByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation meal.GetMealByID has not yet been implemented")
		}),
		MealGetTagsHandler: meal.GetTagsHandlerFunc(func(params meal.GetTagsParams) middleware.Responder {
			return middleware.NotImplemented("operation meal.GetTags has not yet been implemented")
		}),
		InfoGetVersionHandler: info.GetVersionHandlerFunc(func(params info.GetVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation info.GetVersion has not yet been implemented")
		}),
		UserLoginHandler: user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user.Login has not yet been implemented")
		}),
		MealUpdateMealHandler: meal.UpdateMealHandlerFunc(func(params meal.UpdateMealParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation meal.UpdateMeal has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		JWTAuth: func(token string) (*models.User, error) {
			return nil, errors.NotImplemented("api key auth (JWT) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*OpenAPIFoodstoreAPI This is demo for a foodstore (meals) service */
type OpenAPIFoodstoreAPI struct {
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

	// JWTAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	JWTAuth func(string) (*models.User, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// MealCreateMealHandler sets the operation handler for the create meal operation
	MealCreateMealHandler meal.CreateMealHandler
	// MealDeleteMealHandler sets the operation handler for the delete meal operation
	MealDeleteMealHandler meal.DeleteMealHandler
	// MealFindMealsByTagHandler sets the operation handler for the find meals by tag operation
	MealFindMealsByTagHandler meal.FindMealsByTagHandler
	// MealGetIngredientsHandler sets the operation handler for the get ingredients operation
	MealGetIngredientsHandler meal.GetIngredientsHandler
	// InfoGetLivezHandler sets the operation handler for the get livez operation
	InfoGetLivezHandler info.GetLivezHandler
	// MealGetMealByIDHandler sets the operation handler for the get meal by Id operation
	MealGetMealByIDHandler meal.GetMealByIDHandler
	// MealGetTagsHandler sets the operation handler for the get tags operation
	MealGetTagsHandler meal.GetTagsHandler
	// InfoGetVersionHandler sets the operation handler for the get version operation
	InfoGetVersionHandler info.GetVersionHandler
	// UserLoginHandler sets the operation handler for the login operation
	UserLoginHandler user.LoginHandler
	// MealUpdateMealHandler sets the operation handler for the update meal operation
	MealUpdateMealHandler meal.UpdateMealHandler

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
func (o *OpenAPIFoodstoreAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *OpenAPIFoodstoreAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *OpenAPIFoodstoreAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *OpenAPIFoodstoreAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *OpenAPIFoodstoreAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *OpenAPIFoodstoreAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *OpenAPIFoodstoreAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *OpenAPIFoodstoreAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *OpenAPIFoodstoreAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the OpenAPIFoodstoreAPI
func (o *OpenAPIFoodstoreAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.JWTAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.MealCreateMealHandler == nil {
		unregistered = append(unregistered, "meal.CreateMealHandler")
	}
	if o.MealDeleteMealHandler == nil {
		unregistered = append(unregistered, "meal.DeleteMealHandler")
	}
	if o.MealFindMealsByTagHandler == nil {
		unregistered = append(unregistered, "meal.FindMealsByTagHandler")
	}
	if o.MealGetIngredientsHandler == nil {
		unregistered = append(unregistered, "meal.GetIngredientsHandler")
	}
	if o.InfoGetLivezHandler == nil {
		unregistered = append(unregistered, "info.GetLivezHandler")
	}
	if o.MealGetMealByIDHandler == nil {
		unregistered = append(unregistered, "meal.GetMealByIDHandler")
	}
	if o.MealGetTagsHandler == nil {
		unregistered = append(unregistered, "meal.GetTagsHandler")
	}
	if o.InfoGetVersionHandler == nil {
		unregistered = append(unregistered, "info.GetVersionHandler")
	}
	if o.UserLoginHandler == nil {
		unregistered = append(unregistered, "user.LoginHandler")
	}
	if o.MealUpdateMealHandler == nil {
		unregistered = append(unregistered, "meal.UpdateMealHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *OpenAPIFoodstoreAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *OpenAPIFoodstoreAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "JWT":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.JWTAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *OpenAPIFoodstoreAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *OpenAPIFoodstoreAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *OpenAPIFoodstoreAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *OpenAPIFoodstoreAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the open API foodstore API
func (o *OpenAPIFoodstoreAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *OpenAPIFoodstoreAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/meal/{id}"] = meal.NewCreateMeal(o.context, o.MealCreateMealHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/meal/{id}"] = meal.NewDeleteMeal(o.context, o.MealDeleteMealHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/meal/findByTag"] = meal.NewFindMealsByTag(o.context, o.MealFindMealsByTagHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ingredients"] = meal.NewGetIngredients(o.context, o.MealGetIngredientsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/livez"] = info.NewGetLivez(o.context, o.InfoGetLivezHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/meal/{id}"] = meal.NewGetMealByID(o.context, o.MealGetMealByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tags"] = meal.NewGetTags(o.context, o.MealGetTagsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/version"] = info.NewGetVersion(o.context, o.InfoGetVersionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = user.NewLogin(o.context, o.UserLoginHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/meal/{id}"] = meal.NewUpdateMeal(o.context, o.MealUpdateMealHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *OpenAPIFoodstoreAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *OpenAPIFoodstoreAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *OpenAPIFoodstoreAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *OpenAPIFoodstoreAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *OpenAPIFoodstoreAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
