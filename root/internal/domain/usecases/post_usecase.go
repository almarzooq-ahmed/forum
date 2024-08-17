package usecases

import (
	"context"

	"forum/root/internal/domain/entities"
	"forum/root/internal/domain/repositories"
	db_models "forum/root/internal/models/db"
	request_models "forum/root/internal/models/requests"
)

type PostUseCase interface {
	CreatePost(ctx context.Context, req *request_models.CreatePostRequest) (*entities.Post, error)
	GetPostByID(ctx context.Context, id uint) (*entities.Post, error)
	UpdatePost(ctx context.Context, req *request_models.UpdatePostRequest) (*entities.Post, error)
	DeletePost(ctx context.Context, id uint) error
}

type postUseCase struct {
	postRepo repositories.PostRepository
}

func NewPostUseCase(postRepo repositories.PostRepository) PostUseCase {
	return &postUseCase{postRepo: postRepo}
}

func (uc *postUseCase) CreatePost(ctx context.Context, req *request_models.CreatePostRequest) (*entities.Post, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	post := &db_models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := uc.postRepo.Create(ctx, post); err != nil {
		return nil, err
	}

	return &entities.Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}, nil
}

func (uc *postUseCase) GetPostByID(ctx context.Context, id uint) (*entities.Post, error) {
	post, err := uc.postRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entities.Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}, nil
}

func (uc *postUseCase) UpdatePost(ctx context.Context, req *request_models.UpdatePostRequest) (*entities.Post, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	post, err := uc.postRepo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	post.Title = req.Title
	post.Content = req.Content

	if err := uc.postRepo.Update(ctx, post); err != nil {
		return nil, err
	}

	return &entities.Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}, nil
}

func (uc *postUseCase) DeletePost(ctx context.Context, id uint) error {
	return uc.postRepo.Delete(ctx, id)
}
