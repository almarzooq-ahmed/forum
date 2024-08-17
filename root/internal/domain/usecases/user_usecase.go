package usecases

import (
	"context"
	"errors"
	"time"

	entities "forum/root/internal/domain/entities"
	repositories "forum/root/internal/domain/repositories"
	db_models "forum/root/internal/models/db"
	request_models "forum/root/internal/models/requests"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	RegisterUser(ctx context.Context, req *request_models.RegisterUserRequest) (*entities.User, error)
	LoginUser(ctx context.Context, req *request_models.LoginUserRequest) (string, error)
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
}

type userUseCase struct {
	userRepo repositories.UserRepository
	jwtKey   string
}

func NewUserUseCase(userRepo repositories.UserRepository, jwtKey string) UserUseCase {
	return &userUseCase{userRepo: userRepo, jwtKey: jwtKey}
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

func (uc *userUseCase) LoginUser(ctx context.Context, req *request_models.LoginUserRequest) (string, error) {
	user, err := uc.userRepo.FindByUsername(req.Username)
	if err != nil || user == nil || !checkPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := generateJWT(user.Username, uc.jwtKey)
	if err != nil {
		return "", err
	}

	return token, nil
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

// Password hashing using bycrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Compares 2 passwords hash
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// generates a jwt token
func generateJWT(username, jwtKey string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}
