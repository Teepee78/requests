package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/kompere/kompere-api/api/middlewares"
)

func StartServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// CORS middleware
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum age for preflight requests
	})

	// Use CORS middleware
	r.Use(cors.Handler)

	// custom middlewares
	r.Use(middlewares.ParseBody)

	// routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Running!"))
	})
	r.Post("/search", search)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("Method is not valid"))
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		panic("PORT must be set")
	}

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
}
