package di

import (
	echo_logger "GOLANG-DDD-MONGO-TEMPLATE/shared/log/echo_logger"
	"GOLANG-DDD-MONGO-TEMPLATE/shared/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
)

func NewEcho(logger *echo_logger.LoggerWrapper, validator validator.Validator) (*echo.Echo, error) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Logger = logger
	e.Validator = validator
	return e, nil
}

func init() {
	if err := Container.Provide(NewEcho); err != nil {
		panic(errors.Wrap(err, "failed to provide echo"))
	}
}
