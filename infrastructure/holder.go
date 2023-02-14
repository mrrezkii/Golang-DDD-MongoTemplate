package infrastructure

import (
	"GOLANG-DDD-MONGO-TEMPLATE/infrastructure/http/example/example_controller"
	"GOLANG-DDD-MONGO-TEMPLATE/interfaces"
	"GOLANG-DDD-MONGO-TEMPLATE/shared"
	"fmt"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In

		// - example-http-start
		ExampleController *example_controller.Controller
		// - example-http-end

		SharedHolder    shared.Holder
		InterfaceHolder interfaces.Holder
	}
)

func (h *Holder) ListenHTTP() {
	RegisterMiddleware(h)

	h.SharedHolder.Echo.GET("/test", h.ExampleController.FindAll)

	if err := h.SharedHolder.Echo.Start(fmt.Sprintf(":%d", h.SharedHolder.Config.EchoServerPort)); err != nil {
		if err.Error() == "http: Server closed" {
			h.SharedHolder.Logger.Info("closing echo http server")
		} else {
			h.SharedHolder.Logger.Errorf("failed to start echo http server %s", err)
		}
	}
}
