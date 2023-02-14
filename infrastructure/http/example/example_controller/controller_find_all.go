package example_controller

import (
	"GOLANG-DDD-MONGO-TEMPLATE/shared/dto/example"
	"GOLANG-DDD-MONGO-TEMPLATE/shared/util"
	utilctx "GOLANG-DDD-MONGO-TEMPLATE/shared/util/context"
	"github.com/labstack/echo"
)

func (c *Controller) FindAll(ec echo.Context) error {
	var (
		ctx     = ec.Request().Context()
		request example.ExampleRequestDto
	)

	request.MandatoryRequest = utilctx.MandatoryRequest(ctx)

	//response, _ := c.InterfaceHolder.ExampleService.FindAll(ctx, request)
	return util.Response(ec, request.MandatoryRequest, nil)
	//
	//
	//if err != nil {
	//	return util.Response(ec, nil, tasha_error.New(tasha_error.BAD_REQUEST, err))
	//}
	//
	//return util.Response(ec, response.RequestID, nil)
}
