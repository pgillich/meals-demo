package internal

/*
import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/swag"
	flags "github.com/jessevdk/go-flags"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal/restapi"
	"github.com/pgillich/meals-demo/internal/restapi/operations"
)
*/
// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	flags "github.com/jessevdk/go-flags"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal/api"
)

type ServiceFactory func(config configs.Options) api.StrictServerInterface

type ServerOptions struct {
	Host string `long:"host" description:"the IP to listen on" default:"localhost" env:"HOST"`
	Port int    `long:"port" description:"the port to listen on for insecure connections, defaults to a random value" env:"PORT"`
}

func BuildServer(serviceOptions *configs.Options, serviceFactory ServiceFactory) *http.Server {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	serverOptions := &ServerOptions{}
	parser := flags.NewParser(serverOptions, flags.Default)
	parser.ShortDescription = "OpenAPI Foodstore"
	parser.LongDescription = "This is demo for a foodstore (meals) service"
	//serverOptions.ConfigureFlags()
	_, err = parser.AddGroup("service", "Service options", serviceOptions)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	foodStore := serviceFactory(*serviceOptions)

	foodStoreStrictHandler := api.NewStrictHandler(foodStore, nil)

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(context.Context, *openapi3filter.AuthenticationInput) error { return nil },
		},
	}))

	// We now register our petStore above as the handler for the interface
	api.HandlerFromMux(foodStoreStrictHandler, r)

	s := &http.Server{ //_nolint:gosec // Slowloris not important
		Handler: r,
		Addr:    fmt.Sprintf("%s:%d", serverOptions.Host, serverOptions.Port),
	}

	// And we serve HTTP until the world ends.
	//log.Fatal(s.ListenAndServe())
	return s
}
