package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the configuration for the client.
type Config struct {
	ServerAddress string
}

// LoadConfig loads configuration from environment variables.
// It also loads the .env file if present.
func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		return nil, fmt.Errorf("SERVER_ADDRESS environment variable is required")
	}

	return &Config{
		ServerAddress: serverAddress,
	}, nil
}
