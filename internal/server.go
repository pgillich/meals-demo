package internal

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

func BuildServer(
	serviceOptions *configs.Options,
	serviceApis ...func(configs.Options, *operations.OpenAPIFoodstoreAPI),
) *restapi.Server {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatal(err)
	}

	api := operations.NewOpenAPIFoodstoreAPI(swaggerSpec)

	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, swag.CommandLineOptionsGroup{
		ShortDescription: "service",
		LongDescription:  "Service options",
		Options:          serviceOptions,
	})

	server := restapi.NewServer(api)

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "OpenAPI Foodstore"
	parser.LongDescription = "This is demo for a foodstore (meals) service"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatal(err)
		}
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

	for _, serviceAPI := range serviceApis {
		serviceAPI(*serviceOptions, api)
	}

	server.ConfigureAPI()

	return server
}
