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
	// inject the session loading and saving middleware, the next handler will be called after loading and before saving so the next handler can modify the request sessions before saving it
	mux.Use(app.SessionLoad)

	// register endpoints to be handlerd
	mux.Get("/", app.HomeHandler)
	mux.Post("/login", app.LoginHandler)
	mux.Post("/logout", app.LogoutHandler)
	mux.Post("/register", app.RegisterHandler)
	mux.Post("/buy-subscription", app.BuySubscribtionHandler)
	mux.Post("/activate-account", app.ActivateAccountHandler)

	// return the defined router handler mux
	return mux
}
