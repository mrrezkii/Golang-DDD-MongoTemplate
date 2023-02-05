package shared

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/config"
	logger "SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/log"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"io"
	"log"
)

type (
	Holder struct {
		Config *config.EnvConfiguration

		Echo *echo.Echo

		Logger logger.LoggerWithContext

		Validator validator.Validate
	}
)

func (h *Holder) Close() {
	log.Println("closing resource")

	var i interface{} = nil

	i = h.Echo
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
}
