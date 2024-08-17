package request_models

import "github.com/go-playground/validator/v10"

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginUserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
