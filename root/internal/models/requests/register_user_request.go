package request_models

import "github.com/go-playground/validator/v10"

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *RegisterUserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
