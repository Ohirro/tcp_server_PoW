package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
}

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
