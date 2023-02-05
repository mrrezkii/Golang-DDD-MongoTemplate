package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"os"
)

type (
	EnvConfiguration struct {

		// - echo-config-start

		EchoServerPort int `envconfig:"ECHO_SERVER_PORT" default:"9000"`

		// - echo-config-end

		// - logger-config-start

		LoggerFileName string `envconfig:"LOGGER_FILE_NAME"`

		LoggerFormater string `envconfig:"LOGGER_FORMATTER" default:"TEXT" required:"true"`

		LoggerLevel string `envconfig:"LOGGER_LEVEL" default:"INFO" required:"true"`

		LoggerMaxSize int `envconfig:"LOGGER_MAX_SIZE" default:"400" required:"true"`

		LoggerMaxBackups int `envconfig:"LOGGER_MAX_BACKUPS" default:"0" required:"true"`

		LoggerMaxAge int `envconfig:"LOGGER_MAX_AGE" default:"7" required:"true"`

		LoggerCompress bool `envconfig:"LOGGER_COMPRESS" default:"true" required:"true"`

		// - logger-config-end
	}
)

func NewEnvConfiguration() (*EnvConfiguration, error) {
	configuration := EnvConfiguration{}

	if err := NewConfig(&configuration); err != nil {
		return nil, errors.Wrap(err, "failed to provide env config")
	}

	return &configuration, nil
}

func NewConfig(object interface{}) error {
	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", object); err != nil {
			return errors.Wrap(err, "failed to read from env variable")
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", object); err != nil {
		return errors.Wrap(err, "failed to read from env variable")
	}

	return nil
}
