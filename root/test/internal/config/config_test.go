package config_test

import (
	"os"
	"testing"

	"forum/root/internal/config"
)

// Helper function to set up environment for tests
func setEnvVars(t *testing.T, vars map[string]string) func() {
	for key, value := range vars {
		err := os.Setenv(key, value)
		if err != nil {
			t.Fatalf("Failed to set environment variable %s: %v", key, err)
		}
	}
	return func() {
		for key := range vars {
			os.Unsetenv(key)
		}
	}
}

func TestLoadConfig_Success(t *testing.T) {
	// Arrange
	restore := setEnvVars(t, map[string]string{
		"DATABASE_DSN": "user:password@/dbname",
		"SERVER_PORT":  "8080",
	})
	defer restore()

	// Act
	config, err := config.LoadConfig()
	// Assert
	if err != nil {
		t.Fatalf("LoadConfig() returned an error: %v", err)
	}

	if config.DatabaseDSN != "user:password@/dbname" {
		t.Errorf("Expected DATABASE_DSN to be 'user:password@/dbname', got '%s'", config.DatabaseDSN)
	}

	if config.ServerPort != 8080 {
		t.Errorf("Expected SERVER_PORT to be 8080, got %d", config.ServerPort)
	}
}

func TestLoadConfig_MissingDatabaseDSN(t *testing.T) {
	// Arrange
	restore := setEnvVars(t, map[string]string{
		"SERVER_PORT": "8080",
	})
	defer restore()

	// Act
	_, err := config.LoadConfig()

	// Assert
	if err == nil {
		t.Fatal("Expected an error, but got nil")
	}

	if err.Error() != "DATABASE_DSN environment variable is required" {
		t.Errorf("Expected error message 'DATABASE_DSN environment variable is required', got '%v'", err)
	}
}

func TestLoadConfig_InvalidServerPort(t *testing.T) {
	// Arrange
	restore := setEnvVars(t, map[string]string{
		"DATABASE_DSN": "user:password@/dbname",
		"SERVER_PORT":  "invalid-port",
	})
	defer restore()

	// Act
	config, err := config.LoadConfig()
	// Assert
	if err != nil {
		t.Fatalf("LoadConfig() returned an error: %v", err)
	}

	if config.ServerPort != 8080 {
		t.Errorf("Expected SERVER_PORT to default to 8080, got %d", config.ServerPort)
	}
}

func TestLoadConfig_EmptyServerPort(t *testing.T) {
	// Arrange
	restore := setEnvVars(t, map[string]string{
		"DATABASE_DSN": "user:password@/dbname",
		"SERVER_PORT":  "",
	})
	defer restore()

	// Act
	config, err := config.LoadConfig()
	// Assert
	if err != nil {
		t.Fatalf("LoadConfig() returned an error: %v", err)
	}

	if config.ServerPort != 8080 {
		t.Errorf("Expected SERVER_PORT to default to 8080, got %d", config.ServerPort)
	}
}
