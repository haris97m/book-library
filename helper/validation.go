package helper

import (
	"github.com/go-playground/validator/v10"
)

// STRUCTURE ERROR VALIDATION
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New() //call validator package

// Validate Structure of Request
func ValidateStruct(request interface{}) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
