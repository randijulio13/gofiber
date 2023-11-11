package request

import (
	"github.com/randijulio13/gofiber/pkg/validator"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=8"`
}

func (request *LoginRequest) Validate() validator.ValidationErrors {
	return validator.ValidateStruct(request)
}
