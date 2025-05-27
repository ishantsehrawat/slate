package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnv returns the value of the environment variable named by the key.
// If the variable is empty or not present, it returns the specified fallback value.
func GetEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists && value != "" {
        return value
    }
    return fallback
}

// LoadEnv loads environment variables from a .env file if it exists.
// If .env file is not found, it silently proceeds, assuming variables are set elsewhere.
func LoadEnv() {
    // Only load .env if running in non-production (optional)
    env := os.Getenv("GO_ENV")

    if env == "production" {
        return
    }

    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found or failed to load, proceeding with system environment variables")
    } else {
        log.Println(".env file loaded successfully")
    }
}