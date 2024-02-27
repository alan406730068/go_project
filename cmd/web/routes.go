package main

import (
	"net/http"

	"github.com/alan406730068/go_project/pkg/config"
	"github.com/alan406730068/go_project/pkg/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.Appconfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSruf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
