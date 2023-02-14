package di

import (
	"GOLANG-DDD-MONGO-TEMPLATE/shared/config"
	logs "GOLANG-DDD-MONGO-TEMPLATE/shared/log/logs"
	"github.com/pkg/errors"
)

func NewLogger(conf *config.EnvConfiguration) (logs.Logger, error) {
	return logs.New(&logs.Option{
		LogFilePath: conf.LoggerFileName,
		Formatter:   logs.Formatter(conf.LoggerFormater),
		Level:       logs.Level(conf.LoggerLevel),
		MaxSize:     conf.LoggerMaxSize,
		MaxBackups:  conf.LoggerMaxBackups,
		MaxAge:      conf.LoggerMaxAge,
		Compress:    conf.LoggerCompress,
	})
}

func init() {
	if err := Container.Provide(NewLogger); err != nil {
		panic(errors.Wrap(err, "failed to provide logger"))
	}
}
