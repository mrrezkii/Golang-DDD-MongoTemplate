package application

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/application/example"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In

		// - example-domain-start
		ExampleService example.Service
		// - example-domain-end
	}
)
