package database

import (
	db_models "forum/root/internal/models/db"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Creates sqlite databse
func CreateDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = migrateTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

// Creates tables migrations
func migrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&db_models.User{},
		&db_models.Post{},
	)
}
