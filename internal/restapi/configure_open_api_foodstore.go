// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/pgillich/meals-demo/internal/models"
	"github.com/pgillich/meals-demo/internal/restapi/operations"
	"github.com/pgillich/meals-demo/internal/restapi/operations/info"
	"github.com/pgillich/meals-demo/internal/restapi/operations/meal"
	"github.com/pgillich/meals-demo/internal/restapi/operations/user"
)

//go:generate swagger generate server --target ../../internal --name OpenAPIFoodstore --spec ../../api/foodstore.yaml --principal models.User --exclude-main

func configureFlags(api *operations.OpenAPIFoodstoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OpenAPIFoodstoreAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.JWTAuth == nil {
		api.JWTAuth = func(token string) (*models.User, error) {
			return nil, errors.NotImplemented("api key auth (JWT) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.MealCreateMealHandler == nil {
		api.MealCreateMealHandler = meal.CreateMealHandlerFunc(func(params meal.CreateMealParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation meal.CreateMeal has not yet been implemented")
		})
	}
	if api.MealDeleteMealHandler == nil {
		api.MealDeleteMealHandler = meal.DeleteMealHandlerFunc(func(params meal.DeleteMealParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation meal.DeleteMeal has not yet been implemented")
		})
	}
	if api.MealFindMealsByTagHandler == nil {
		api.MealFindMealsByTagHandler = meal.FindMealsByTagHandlerFunc(func(params meal.FindMealsByTagParams) middleware.Responder {
			return middleware.NotImplemented("operation meal.FindMealsByTag has not yet been implemented")
		})
	}
	if api.MealGetIngredientsHandler == nil {
		api.MealGetIngredientsHandler = meal.GetIngredientsHandlerFunc(func(params meal.GetIngredientsParams) middleware.Responder {
			return middleware.NotImplemented("operation meal.GetIngredients has not yet been implemented")
		})
	}
	if api.InfoGetLivezHandler == nil {
		api.InfoGetLivezHandler = info.GetLivezHandlerFunc(func(params info.GetLivezParams) middleware.Responder {
			return middleware.NotImplemented("operation info.GetLivez has not yet been implemented")
		})
	}
	if api.MealGetMealByIDHandler == nil {
		api.MealGetMealByIDHandler = meal.GetMealByIDHandlerFunc(func(params meal.GetMealByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation meal.GetMealByID has not yet been implemented")
		})
	}
	if api.MealGetTagsHandler == nil {
		api.MealGetTagsHandler = meal.GetTagsHandlerFunc(func(params meal.GetTagsParams) middleware.Responder {
			return middleware.NotImplemented("operation meal.GetTags has not yet been implemented")
		})
	}
	if api.InfoGetVersionHandler == nil {
		api.InfoGetVersionHandler = info.GetVersionHandlerFunc(func(params info.GetVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation info.GetVersion has not yet been implemented")
		})
	}
	if api.UserLoginHandler == nil {
		api.UserLoginHandler = user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user.Login has not yet been implemented")
		})
	}
	if api.MealUpdateMealHandler == nil {
		api.MealUpdateMealHandler = meal.UpdateMealHandlerFunc(func(params meal.UpdateMealParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation meal.UpdateMeal has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
