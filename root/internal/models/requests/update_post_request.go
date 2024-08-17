package request_models

import "github.com/go-playground/validator/v10"

type UpdatePostRequest struct {
	ID      uint   `json:"id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (r *UpdatePostRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
