package config

import "os"

type Config struct {
	DatabaseURL string
	JWTSecret   string
	AppPort     string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://desent:desent123@localhost:5432/desent_pretest?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "desent-pretest-secret-key-2024"),
		AppPort:     getEnv("PORT", getEnv("APP_PORT", "3000")),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
