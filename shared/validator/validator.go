package validator

import "github.com/go-playground/validator"

type (
	Validator interface {
		Validate(interface{}) error
		ValidateVar(interface{}, string) error
	}

	validate struct {
		instance *validator.Validate
	}
)

func New() Validator {
	return &validate{instance: validator.New()}
}

func (v *validate) Validate(object interface{}) error {
	if err := v.instance.Struct(object); err != nil {
		return err
	}

	return nil
}

func (v *validate) ValidateVar(object interface{}, constraint string) error {
	if err := v.instance.Var(object, constraint); err != nil {
		return err
	}

	return nil
}
