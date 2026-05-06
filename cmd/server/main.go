package main

import (
	"net/http"

	"github.com/dborah123/fhir-ingestion-service/internal/api"
	"github.com/dborah123/fhir-ingestion-service/internal/publisher"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// Standard Middleware
	r.Use(middleware.Logger)    // Structured logging
	r.Use(middleware.Recoverer) // Don't crash on panics

	// Dependencies
	pub := publisher.NewMockPublisher()

	// Routes
	r.Get("/health", api.HealthHandler(pub))
	r.Post("/ingest/fhir", api.FhirIngest())

	http.ListenAndServe(":8080", r)
}
