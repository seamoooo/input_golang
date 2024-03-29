package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)

	// ipアドレスの直接リクエストでもいけるようにする
	mux.Use(app.appIpToContext)

	// register routes
	mux.Get("/", app.Home)

	// static assets
	fileServer := http.FileServer(http.Dir("./static/assets"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
