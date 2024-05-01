package common

import (
	"github.com/joho/godotenv"
	"os"
)

// LoadEnv
// Loading .env file from passed argument (if present) or default location
func LoadEnv() {
	if len(os.Args) > 1 {

		envPath := os.Args[1]

		if _, err := os.Stat(envPath); err == nil {
			_ = godotenv.Load(envPath)
		} else {
			_ = godotenv.Load()
		}
	} else {
		_ = godotenv.Load()
	}
}

// GetEnv
// Get environment variable, if not present returns fallback value
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
