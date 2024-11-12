package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	//posts

	//users

	//auth

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 8,
		IdleTimeout:  time.Minute * 1,
	}
	log.Printf("Server is statred at %s", app.config.addr)
	return srv.ListenAndServe()
}
