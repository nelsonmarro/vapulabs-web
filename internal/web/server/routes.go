// Package server defines the HTTP routes for the web server.
package server

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nelsonmarro/vapulabs-web/internal/web/handlers"
)

func (s *Server) getRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	// Static files
	staticDir := os.Getenv("STATIC_DIR")
	if staticDir == "" {
		staticDir = "./static"
	}
	fs := http.FileServer(http.Dir(staticDir))
	mux.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Handlers
	homeHandler := handlers.NewHomeHandler()
	mux.Get("/", homeHandler.ServeHTTP)

	productHandler := handlers.NewProductHandler()
	mux.Get("/products/{id}", productHandler.ServeHTTP)
	mux.Get("/products/{id}/pricing", productHandler.ServePricing)
	mux.Get("/products/{id}/pricing/view", productHandler.ServePricingGrid)
	mux.Get("/products/{id}/download", productHandler.ServeDownload)

	contactHandler := s.createContactHandler()
	mux.Post("/contact", contactHandler.HandleSubmit)
	mux.Get("/contact/form", contactHandler.ServeForm)

	legalHandler := handlers.NewLegalHandler()
	mux.Get("/privacy", legalHandler.ServePrivacy)
	mux.Get("/legal", legalHandler.ServeTerms)

	return mux
}
