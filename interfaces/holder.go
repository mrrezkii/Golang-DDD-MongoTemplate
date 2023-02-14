package interfaces

import (
	"GOLANG-DDD-MONGO-TEMPLATE/interfaces/example"
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
