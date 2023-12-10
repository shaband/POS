package helper

import (
	"github.com/go-playground/validator"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		
		validator *validator.Validate
	}
)

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func RegisterValidator(v *validator.Validate) *XValidator {
	return &XValidator{
		validator: v,
	}
}
