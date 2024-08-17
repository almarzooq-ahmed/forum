package repositories

import (
	"context"

	db_models "forum/root/internal/models/db"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *db_models.User) error
	FindByUsername(username string) (*db_models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *db_models.User) error {
	result := r.db.WithContext(ctx).Create(user)
	return result.Error
}

func (r *userRepository) FindByUsername(username string) (*db_models.User, error) {
	var user db_models.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
