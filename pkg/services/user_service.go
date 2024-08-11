package services

import (
	"context"

	"forum/pkg/entities"
	"forum/pkg/models"
	"forum/pkg/repositories"
)

type UserService interface {
	RegisterUser(ctx context.Context, username, email, password string) (*entities.UserEntity, error)
	GetUserByID(ctx context.Context, id uint) (*entities.UserEntity, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(ctx context.Context, username, email, password string) (*entities.UserEntity, error) {
	userModel := &models.UserModel{
		Username: username,
		Email:    email,
		Password: password,
	}

	createdUser, err := s.userRepo.Create(ctx, userModel)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.UserEntity{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}

	return userEntity, nil
}

func (s *userService) GetUserByID(ctx context.Context, id uint) (*entities.UserEntity, error) {
	userModel, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.UserEntity{
		ID:       userModel.ID,
		Username: userModel.Username,
		Email:    userModel.Email,
	}

	return userEntity, nil
}
