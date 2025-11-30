package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func TranslateErrorMessage(err error) map[string]string {
	// initiate map for error messages
	errorsMap := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			fieldName := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				errorsMap[strings.ToLower(fieldName)] = fmt.Sprintf("%s is required", fieldName)
			case "email":
				errorsMap[strings.ToLower(fieldName)] = "Invalid email format"
			case "unique":
				errorsMap[strings.ToLower(fieldName)] = fmt.Sprintf("%s must be unique", fieldName)
			case "min":
				errorsMap[strings.ToLower(fieldName)] = fmt.Sprintf("%s must be at least %s characters long", fieldName, fieldError.Param())
			case "max":
				errorsMap[strings.ToLower(fieldName)] = fmt.Sprintf("%s must be at most %s characters long", fieldName, fieldError.Param())
			case "numeric":
				errorsMap[strings.ToLower(fieldName)] = fmt.Sprintf("%s must be a numeric value", fieldName)
			default:
				errorsMap[strings.ToLower(fieldName)] = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fieldName, fieldError.Tag())
			}
		}
	}

	if err != nil {
		if strings.Contains(err.Error(), "Duplicated entry") {
			if strings.Contains(err.Error(), "username") {
				errorsMap["username"] = "Username already exists"
			}
			if strings.Contains(err.Error(), "email") {
				errorsMap["email"] = "Email already exists"
			}
		}
	}

	return errorsMap
}

func IsDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicated entry")
}
