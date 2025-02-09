package main

import (
	"log"

	"github.com/swaniket/social/internal/env"
)

func main() {
	configuration := config{
		addr: env.GetEnvVariableAsString("ADDR", ":8080"),
	}

	app := &application{
		serverConfig: configuration,
	}

	mux := app.mountApplication()
	log.Fatal(app.runApplication(mux))
}
