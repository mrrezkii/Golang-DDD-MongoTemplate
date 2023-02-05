package di

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/application"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/infrastructure"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/interfaces"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

var (
	Container *dig.Container = dig.New()
)

func init() {
	if err := application.Register(Container); err != nil {
		panic(errors.Wrap(err, "failed to register application container"))
	}

	if err := interfaces.Register(Container); err != nil {
		panic(errors.Wrap(err, "failed to register interfaces container"))
	}

	if err := infrastructure.Register(Container); err != nil {
		panic(errors.Wrap(err, "failed to register infrastructure container"))
	}

	if err := shared.Register(Container); err != nil {
		panic(errors.Wrap(err, "failed to register shared container"))
	}
}
