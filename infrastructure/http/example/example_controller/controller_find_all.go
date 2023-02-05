package example_controller

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/dto/example"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/util"
	utilctx "SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/util/context"
	tasha_error "SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/util/error"
	"github.com/labstack/echo"
)

func (c *ControllerImpl) FindAll(ec echo.Context) error {
	var (
		ctx     = ec.Request().Context()
		request example.ExampleRequestDto
	)

	request.MandatoryRequest = utilctx.MandatoryRequest(ctx)

	response, err := c.ApplicationHolder.ExampleService.FindAll(ctx, request)
	
	if err != nil {
		return util.Response(ec, nil, tasha_error.New(tasha_error.BAD_REQUEST, err))
	}

	return util.Response(ec, response.RequestID, nil)
}
