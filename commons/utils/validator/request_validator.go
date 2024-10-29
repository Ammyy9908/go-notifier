package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

type IRequestValidator interface {
	Validate(i interface{}) error
}

func (v *Validator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}
