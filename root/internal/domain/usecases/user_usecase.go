package usecases

import (
	"context"

	entities "forum/root/internal/domain/entities"
	repositories "forum/root/internal/domain/repositories"
)

type UserUseCase interface {
	GetUserById(ctx context.Context, id string) (*entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
}

type userUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (uc *userUseCase) GetUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	user, err := uc.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (uc *userUseCase) GetUserById(ctx context.Context, id string) (*entities.User, error) {
	user, err := uc.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
