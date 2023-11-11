package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type ValidationError struct {
	Key     string `json:"key"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

type ValidationErrors []ValidationError

func InitializeValidator() *validator.Validate {
	validate = validator.New()
	registerJsonTag()
	return validate
}

func registerJsonTag() {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func ValidateStruct(request interface{}) ValidationErrors {
	if err := validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorsField ValidationErrors
		for _, fieldError := range validationErrors {
			errorsField = append(errorsField, ValidationError{
				Key:     fieldError.Field(),
				Tag:     fieldError.Tag(),
				Message: GetValidationErrorMessage(fieldError),
			})
		}
		return errorsField
	}
	return nil
}

func GetValidator() *validator.Validate {
	return validate
}
