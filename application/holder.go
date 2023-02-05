package application

import "SANDBOX-TASHA-ISSUER-SERVICE-BE/application/example"

type (
	Holder struct {
		// - example-domain-start
		ExampleService example.Service
		// - example-domain-end
	}
)
