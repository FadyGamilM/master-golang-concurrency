package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// register routes to the server handler and return it
func (app *Config) routes() http.Handler {
	// define a chi router instance
	mux := chi.NewRouter()

	// inject middlewares
	mux.Use(middleware.Recoverer)

	// register endpoints to be handlerd
	mux.Get("/", app.HomeHandler)

	// return the defined router handler mux
	return mux
}
