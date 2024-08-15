package usecases

import (
	"context"

	entities "forum/root/internal/domain/entities"
	repositories "forum/root/internal/domain/repositories"
	db_models "forum/root/internal/models/db"
	request_models "forum/root/internal/models/requests"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, req *request_models.RegisterUserRequest) (*entities.User, error)
}

type userUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (uc *userUseCase) RegisterUser(ctx context.Context, req *request_models.RegisterUserRequest) (*entities.User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &db_models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &entities.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

// Password hashing using bycrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
