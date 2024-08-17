package request_models

import "github.com/go-playground/validator/v10"

type CreatePostRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	UserID  uint   `json:"userId" validate:"required"`
}

func (r *CreatePostRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
