package interfaces

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/interfaces/example"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(example.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide example controller")
	}

	return nil
}
