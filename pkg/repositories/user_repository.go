package repositories

import (
	"context"

	"forum/pkg/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.UserModel) (*models.UserModel, error)
	GetByID(ctx context.Context, id uint) (*models.UserModel, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.UserModel) (*models.UserModel, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*models.UserModel, error) {
	var user models.UserModel
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func hashPassword(password string) (string, error) {
	return password, nil
}
