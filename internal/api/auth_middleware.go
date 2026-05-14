package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/dborah123/fhir-ingestion-service/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("local-dev-secret-change-me")

type contextKey string

const ctxKeyClaims contextKey = "claims"

func AuthMiddleware(cfg config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
				return
			}

			raw := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwt.Parse(raw, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return jwtSecret, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, `{"error":"forbidden"}`, http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), ctxKeyClaims, token.Claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
