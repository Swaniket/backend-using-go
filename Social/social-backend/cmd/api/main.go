package main

import (
	"log"

	"github.com/swaniket/social/internal/env"
	"github.com/swaniket/social/internal/store"
)

func main() {
	configuration := config{
		addr: env.GetEnvVariableAsString("ADDR", ":8080"),
	}

	postgresStore := store.NewPostgresStorage(nil) // @TODO: Pass DB conn over here

	app := &application{
		serverConfig: configuration,
		store:        postgresStore,
	}

	mux := app.mountApplication()
	log.Fatal(app.runApplication(mux))
}
