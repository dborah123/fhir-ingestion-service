package main

import (
	"log"
	"net/http"

	"github.com/dborah123/fhir-ingestion-service/internal/api"
	"github.com/dborah123/fhir-ingestion-service/internal/config"
	"github.com/dborah123/fhir-ingestion-service/internal/publisher"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Only enforce JWT auth when running locally
	// In AWS, API Gateway handles this before requests reach us
	if cfg.Env == "local" {
		r.Use(api.AuthMiddleware(cfg))
	}

	pub := publisher.NewMockPublisher()

	r.Get("/health", api.HealthHandler(pub))
	r.Post("/ingest/fhir", api.FhirIngest())

	log.Printf("starting server on :%s (env=%s)", cfg.Port, cfg.Env)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
