package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// Holds application configurations
type Config struct {
	DatabaseDSN string
	ServerPort  int
}

// Loads configurations from .env file
func LoadConfig() (*Config, error) {
	rootDir, err := filepath.Abs("../")
	if err != nil {
		log.Fatal("Error getting root directory:", err)
	}

	err = godotenv.Load(filepath.Join(rootDir, ".env"))

	dbDSN := os.Getenv("DATABASE_DSN")
	if dbDSN == "" {
		return nil, fmt.Errorf("DATABASE_DSN environment variable is required")
	}

	portStr := os.Getenv("SERVER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 8080
	}

	return &Config{
		DatabaseDSN: dbDSN,
		ServerPort:  port,
	}, nil
}
