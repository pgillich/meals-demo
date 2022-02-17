package cmd

import (
	"log"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal"
	"github.com/pgillich/meals-demo/internal/logic"
)

func RunMeals() {
	options := configs.Options{} //nolint:exhaustivestruct // default values
	server := internal.BuildServer(&options, logic.SetInfoAPI, logic.SetUserAPI, logic.SetMealAPI)

	defer server.Shutdown() //nolint:errcheck // never mind at exit
	if err := server.Serve(); err != nil {
		log.Fatal(err) //noling:gocritic // never mind at exit
	}
}
