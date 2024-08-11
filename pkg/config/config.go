package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DatabaseDSN string
	ServerPort  int
}

func LoadConfig() (*Config, error) {
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
