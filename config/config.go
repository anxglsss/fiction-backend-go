package config

import (
	"os"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	CORSOrigin  string
}

func Load() *Config {
	return &Config{
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "data.db"),
		CORSOrigin:  getEnv("CORS_ORIGIN", "http://localhost:3000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
