package handlers

import (
	"net/http"

	usecases "forum/root/internal/domain/usecases"
	request_models "forum/root/internal/models/requests"
	response_models "forum/root/internal/models/responses"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	registerUserUseCase usecases.RegisterUserUseCase
}

func NewUserHandler(registerUserUseCase usecases.RegisterUserUseCase) *UserHandler {
	return &UserHandler{registerUserUseCase: registerUserUseCase}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req request_models.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.registerUserUseCase.RegisterUser(c.Request.Context(), &req)
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
