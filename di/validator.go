package di

import (
	"GOLANG-DDD-MONGO-TEMPLATE/shared/validator"
	"github.com/pkg/errors"
)

func NewValidator() (validator.Validator, error) {
	return validator.New(), nil
}

func init() {
	if err := Container.Provide(NewValidator); err != nil {
		panic(errors.Wrap(err, "failed to provide validator"))
	}
}
