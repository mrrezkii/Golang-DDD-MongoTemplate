package example_controller

import (
	application "SANDBOX-TASHA-ISSUER-SERVICE-BE/interfaces"
	"github.com/labstack/echo"
)

type (
	Controller interface {
		FindAll(c echo.Context) error
	}

	ControllerImpl struct {
		controller        Controller
		ApplicationHolder application.Holder
	}
)
