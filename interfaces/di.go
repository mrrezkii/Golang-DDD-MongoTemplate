package interfaces

import (
	"GOLANG-DDD-MONGO-TEMPLATE/interfaces/example"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(example.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide example controller")
	}

	return nil
}
