package helper

import (
	valid "github.com/asaskevich/govalidator"
)

type Validator struct{}

func (v Validator) Validate(data interface{}) (bool, error) {
	valid.SetFieldsRequiredByDefault(true)
	return valid.ValidateStruct(data)
}

func RegisterValidator() *Validator {
	return &Validator{}
}
