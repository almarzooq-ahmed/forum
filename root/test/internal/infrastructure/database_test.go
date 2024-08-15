package database_test

import (
	"os"
	"testing"

	database "forum/root/internal/infrastructure/database"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Helper function to create a temporary SQLite database
func createTempDB(t *testing.T) (*gorm.DB, func()) {
	// Create a temporary file for the SQLite database
	tmpFile, err := os.CreateTemp("", "testdb-*.db")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	tmpFileName := tmpFile.Name()
	tmpFile.Close() // Close the file handle so it can be removed later

	// Ensure the file is removed after the test
	t.Cleanup(func() {
		if err := os.Remove(tmpFileName); err != nil {
			t.Errorf("Failed to remove temporary file: %v", err)
		}
	})

	db, err := database.CreateDB(tmpFileName)
	if err != nil {
		t.Fatalf("CreateDB() returned an error: %v", err)
	}

	return db, func() {
		if err := os.Remove(tmpFileName); err != nil {
			t.Errorf("Failed to remove temporary file: %v", err)
		}
	}
}

func TestCreateDB_Success(t *testing.T) {
	// Act
	db, cleanup := createTempDB(t)
	defer cleanup()

	// Assert
	assert.NotNil(t, db, "Expected db to be non-nil")
}

func TestMigrateTables(t *testing.T) {
	// Act
	db, cleanup := createTempDB(t)
	defer cleanup()

	// Define a model for testing
	type TestModel struct {
		ID uint `gorm:"primaryKey"`
	}

	// Use AutoMigrate with the TestModel to ensure the table is created
	if err := db.AutoMigrate(&TestModel{}); err != nil {
		t.Fatalf("AutoMigrate() returned an error: %v", err)
	}

	// Assert that the table exists
	hasTable := db.Migrator().HasTable(&TestModel{})
	assert.True(t, hasTable, "Expected TestModel table to be created")
}
