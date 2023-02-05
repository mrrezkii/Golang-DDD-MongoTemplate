package interfaces

import "SANDBOX-TASHA-ISSUER-SERVICE-BE/interfaces/example"

type (
	Holder struct {
		// - example-domain-start
		ExampleService example.Service
		// - example-domain-end
	}
)
