package application

import (
	"GOLANG-DDD-MONGO-TEMPLATE/application/example"
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
