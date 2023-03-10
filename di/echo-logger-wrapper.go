package di

import (
	echo_logger "GOLANG-DDD-MONGO-TEMPLATE/shared/log/echo_logger"
	"GOLANG-DDD-MONGO-TEMPLATE/shared/log/logs"
	"github.com/pkg/errors"
)

func NewEchoLoggerWrapper(logger logs.Logger) (*echo_logger.LoggerWrapper, error) {
	return echo_logger.NewLoggerWrapper(logger)
}

func init() {
	if err := Container.Provide(NewEchoLoggerWrapper); err != nil {
		panic(errors.Wrap(err, "failed to provide echo-logger-wrapper"))
	}
}
