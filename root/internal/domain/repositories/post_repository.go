package repositories

import (
	"context"

	db_models "forum/root/internal/models/db"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(ctx context.Context, post *db_models.Post) error
	FindByID(ctx context.Context, id uint) (*db_models.Post, error)
	Update(ctx context.Context, post *db_models.Post) error
	Delete(ctx context.Context, id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(ctx context.Context, post *db_models.Post) error {
	result := r.db.WithContext(ctx).Create(post)
	return result.Error
}

func (r *postRepository) FindByID(ctx context.Context, id uint) (*db_models.Post, error) {
	var post db_models.Post
	result := r.db.WithContext(ctx).Preload("User").First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func (r *postRepository) Update(ctx context.Context, post *db_models.Post) error {
	result := r.db.WithContext(ctx).Save(post)
	return result.Error
}

func (r *postRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&db_models.Post{}, id)
	return result.Error
}
