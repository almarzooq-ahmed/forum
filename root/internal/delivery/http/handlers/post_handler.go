package handlers

import (
	"net/http"
	"strconv"

	"forum/root/internal/domain/usecases"
	request_models "forum/root/internal/models/requests"
	response_models "forum/root/internal/models/responses"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postUseCase usecases.PostUseCase
}

func NewPostHandler(postUseCase usecases.PostUseCase) *PostHandler {
	return &PostHandler{postUseCase: postUseCase}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req request_models.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.postUseCase.CreatePost(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response_models.GetPostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	})
}

func (h *PostHandler) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.postUseCase.GetPostByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response_models.GetPostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	})
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	var req request_models.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := h.postUseCase.UpdatePost(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response_models.GetPostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.postUseCase.DeletePost(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
