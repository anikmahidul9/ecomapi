package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	//logger *log.Logger
	//db     *sql.DB
}

// mount the config struct to hold configuration settings for the application
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})
	
	//http.ListenAndServe(":3333", r)

	return r
}

//run the application with the config struct

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}
