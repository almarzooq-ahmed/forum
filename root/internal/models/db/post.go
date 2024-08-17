package db_models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"type:text;not null"`
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`
}
