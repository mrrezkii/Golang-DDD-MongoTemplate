package shared

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/config"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/log/echo_logger"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/log/logs"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/validator"
	"github.com/labstack/echo"
	"go.uber.org/dig"
	"io"
	"log"
)

type (
	Holder struct {
		dig.In

		Config *config.EnvConfiguration

		Echo *echo.Echo

		EchoLoggerWrapper *echo_logger.LoggerWrapper

		Logger logs.Logger

		Validator validator.Validator
	}
)

func (h *Holder) Close() {
	log.Println("closing resource")

	var i interface{} = nil

	i = h.Echo
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}

	i = h.EchoLoggerWrapper
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}

	i = h.Logger
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}

	i = h.Validator
	if closer, ok := i.(io.Closer); ok {
		_ = closer.Close()
	}

	log.Print("done closing resource")
}
