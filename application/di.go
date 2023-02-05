package application

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/application/example"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	// - example-domain-start

	if err := container.Provide(example.NewService); err != nil {
		return errors.Wrap(err, "failed to provide example app service")
	}

	// - example-domain-end

	return nil
}
