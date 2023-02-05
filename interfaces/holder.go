package interfaces

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/interfaces/example"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In

		// - example-domain-start
		ExampleService example.ViewService
		// - example-domain-end
	}
)
