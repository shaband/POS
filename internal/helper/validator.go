package helper

import (
	valid "github.com/asaskevich/govalidator"
)

type Validator struct{}

func (v Validator) Validate(data map[string]interface{}) (bool, error) {
	result, err := valid.ValidateStruct(data)
	if err != nil {
		println("error: " + err.Error())
	}
	return result, err
}

func RegisterValidator() *Validator {
	return &Validator{}
}
