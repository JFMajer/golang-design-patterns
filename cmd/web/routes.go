package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/", app.ShowHome)
	mux.Get("/about", app.ShowAbout)
	mux.Get("/{page}", app.ShowPage)

	return mux
}
