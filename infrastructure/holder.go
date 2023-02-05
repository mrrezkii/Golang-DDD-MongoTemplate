package infrastructure

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/infrastructure/http/example/example_controller"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/interfaces"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared"
	"context"
	"fmt"
)

type (
	Holder struct {
		// - example-http-start
		ExampleController example_controller.Controller
		// - example-http-end

		SharedHolder    shared.Holder
		InterfaceHolder interfaces.Holder
	}
)

func (h *Holder) ListenHTTP() {
	RegisterMiddleware(h)

	h.SharedHolder.Echo.GET("/test", h.ExampleController.FindAll)

	if err := h.SharedHolder.Echo.Start(fmt.Sprint(":%d", h.SharedHolder.Config.EchoServerPort)); err != nil {
		if err.Error() == "http: Server closed" {
			h.SharedHolder.Logger.Info(context.Background(), "closing echo http server")
		} else {
			h.SharedHolder.Logger.Errorf(context.Background(), "failed to start echo http server %s", err)
		}
	}
}
