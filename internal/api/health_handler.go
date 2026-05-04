package api

import (
	"encoding/json"
	"net/http"

	"github.com/dborah123/fhir-ingestion-service/internal/publisher"
)

func HealthHandler(publisher publisher.EventPublisher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// It's good practice to set the Content-Type header first
		w.Header().Set("Content-Type", "application/json")

		err := publisher.Ping(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			// Using a map for quick JSON responses is very common in Go
			json.NewEncoder(w).Encode(map[string]string{
				"status": "unhealthy",
				"error":  err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "up"}`))
	}
}
