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

type AuthUseCase interface {
	RegisterUser(ctx context.Context, req *request_models.RegisterUserRequest) (*entities.User, error)
	LoginUser(ctx context.Context, req *request_models.LoginUserRequest) (*string, error)
}

type authUseCase struct {
	userRepo repositories.UserRepository
	jwtKey   string
}

func NewAuthUseCase(userRepo repositories.UserRepository, jwtKey string) AuthUseCase {
	return &authUseCase{userRepo: userRepo, jwtKey: jwtKey}
}

func (uc *authUseCase) RegisterUser(ctx context.Context, req *request_models.RegisterUserRequest) (*entities.User, error) {
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

func (uc *authUseCase) LoginUser(ctx context.Context, req *request_models.LoginUserRequest) (*string, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := uc.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	if user == nil || !checkPasswordHash(req.Password, user.Password) {
		return nil, errors.New("Invalid username or password")
	}

	token, err := generateJWT(user.ID, user.Username, user.Email, uc.jwtKey)
	if err != nil {
		return nil, err
	}

	return &token, nil
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
func generateJWT(id uint, username, email, jwtKey string) (string, error) {
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}
