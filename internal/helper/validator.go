package helper

import (
	valid "github.com/asaskevich/govalidator"
)

type Validator struct{}

func (v Validator) Validate(data interface{}) (bool, error) {
	valid.SetFieldsRequiredByDefault(true)
	return valid.ValidateStruct(data)
}

func (v Validator) StringMatches(s string, params ...string) bool {
	return valid.StringMatches(s, params...)
}

func RegisterValidator() *Validator {
	return &Validator{}
}
