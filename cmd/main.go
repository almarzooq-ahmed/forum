package main

import (
	"log"

	"forum/pkg/database"
)

func main() {
	// Initialize database
	db, err := database.NewSQLiteDB("./forum.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Get the underlying *sql.DB instance
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	defer sqlDB.Close()
}
