package validator

import (
	"fmt"
	vldtr "github.com/go-playground/validator/v10"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *vldtr.Validate {
	// Create a new validator for a Book model.
	validate := vldtr.New()
	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	if _, ok := err.(vldtr.ValidationErrors); ok {
		for _, err := range err.(vldtr.ValidationErrors) {
			errMsg := msgForTag(err.Tag(), err.Param())
			fields[err.Field()] = errMsg
		}

	}

	return fields
}

func msgForTag(tag string, errParam string) string {
	switch tag {
	case "required":
		return "Field ini diperlukan"
	case "required_if":
		return fmt.Sprintf("Field ini diperlukan jika %s", errParam)
	case "email":
		return "Email tidak valid"
	case "lte":
		return fmt.Sprintf("Max karakter adalah %s", errParam)
	case "gte":
		return fmt.Sprintf("Min karakter adalah %s", errParam)
	}
	return ""
}
