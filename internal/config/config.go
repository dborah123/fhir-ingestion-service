package config

import "os"

type Config struct {
	JWTSecret string
	Port      string
	Env       string
}

func Load() Config {
	return Config{
		JWTSecret: getEnv("JWT_SECRET", "local-dev-secret-change-me"),
		Port:      getEnv("PORT", "8080"),
		Env:       getEnv("ENV", "local"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
