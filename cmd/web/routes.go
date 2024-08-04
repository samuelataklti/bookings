package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/samuelataklti/bookings/pkg/config"
	"github.com/samuelataklti/bookings/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	stripPrefix := http.StripPrefix("/static/", fileServer)
	mux.Handle("/static/*", stripPrefix)

	return mux
}
