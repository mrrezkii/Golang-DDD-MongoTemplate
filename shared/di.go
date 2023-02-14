package shared

import (
	"GOLANG-DDD-MONGO-TEMPLATE/shared/config"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(config.NewEnvConfiguration); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
