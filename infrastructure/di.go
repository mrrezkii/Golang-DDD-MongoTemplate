package infrastructure

import (
	"GOLANG-DDD-MONGO-TEMPLATE/infrastructure/http/example/example_controller"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(example_controller.NewController); err != nil {
		return errors.Wrap(err, "failed to provide example controller")
	}

	return nil
}
