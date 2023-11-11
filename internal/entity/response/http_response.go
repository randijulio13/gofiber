package response

import "github.com/randijulio13/gofiber/pkg/validator"

type HttpResponse struct {
	Message string                     `json:"message,omitempty"`
	Status  bool                       `json:"status"`
	Errors  validator.ValidationErrors `json:"errors,omitempty"`
	Data    interface{}                `json:"data,omitempty"`
}
