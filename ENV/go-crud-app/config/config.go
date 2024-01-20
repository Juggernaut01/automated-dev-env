package config

import (
	"os"
)

// AppConfig holds the application configuration
type AppConfig struct {
	DatabaseURL string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() *AppConfig {
	return &AppConfig{
		DatabaseURL: getEnv("DATABASE_URL", "user=yourusername password=yourpassword dbname=crudappdb sslmode=disable"),
	}
}

// getEnv retrieves the value of an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

