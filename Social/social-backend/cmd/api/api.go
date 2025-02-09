package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	serverConfig config
}

type config struct {
	addr string
	// @TODO: Add db, rate limiter config
}

func (app *application) mountApplication() http.Handler {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Recoverer) // This will recover from any panic
	router.Use(middleware.Logger)    // Logger middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	// Grouping the endpoints by version
	router.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return router
}

func (app *application) runApplication(mux http.Handler) error {
	server := &http.Server{
		Addr:         app.serverConfig.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("The server is running at %s", app.serverConfig.addr)

	return server.ListenAndServe()
}
