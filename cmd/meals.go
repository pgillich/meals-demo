package cmd

import (
	"context"
	"log"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal"
	"github.com/pgillich/meals-demo/internal/logic"
)

func RunMeals() {
	options := configs.Options{}
	server := internal.BuildServer(&options, logic.NewFoodStore)

	defer server.Shutdown(context.Background()) //nolint:errcheck // never mind at exit
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err) //noling:gocritic // never mind at exit
	}
}
