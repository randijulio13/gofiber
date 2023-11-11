package validator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorMessage struct {
	Required string `json:"required"`
	Min      string `json:"min"`
}

func GetValidationErrorMessage(fieldError validator.FieldError) string {
	currentDir, _ := os.Getwd()

	file, err := os.ReadFile(currentDir + "/pkg/validator/message.json")
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Dekode file JSON ke map[string]interface{}
	var messages map[string]interface{}
	err = json.Unmarshal(file, &messages)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// Mendapatkan nilai dari kunci "required"
	message, exists := messages[fieldError.Tag()].(string)
	if !exists {
		return fieldError.Error()
	}

	if fieldError.Param() != "" {
		return fmt.Sprintf(message, fieldError.Field(), fieldError.Param())
	}
	return fmt.Sprintf(message, fieldError.Field())
}
