package handlers

import (
	"net/http"

	usecases "forum/root/internal/domain/usecases"
	response_models "forum/root/internal/models/responses"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	username, ok := claims.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	user, err := h.userUseCase.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response_models.GetUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
}
