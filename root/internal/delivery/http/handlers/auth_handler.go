package handlers

import (
	"net/http"

	usecases "forum/root/internal/domain/usecases"
	request_models "forum/root/internal/models/requests"
	response_models "forum/root/internal/models/responses"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUseCase usecases.AuthUseCase
}

func NewAuthHandler(authUseCase usecases.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req request_models.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authUseCase.RegisterUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response_models.RegisterUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request_models.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUseCase.LoginUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response_models.LoginUserResponse{
		Token: *token,
	})
}
