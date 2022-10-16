package config

import (
	"errors"
	"os"

	"github.com/Miavega/constants"
)

func (c *Config) getEnv() (err error) {
	env := os.Getenv(constants.ENV)

	switch env {
	case constants.LOCAL, constants.DEVELOP, constants.PRODUCTION:
		c.BaseConfig.Environment = env
	case constants.EMPTY:
		c.BaseConfig.Environment = constants.LOCAL
	default:
		err = errors.New("env not valid: " + env)
	}

	return
}

func (c *Config) getServerMode() (err error) {
	switch c.BaseConfig.Environment {
	case constants.LOCAL:
		c.Server.Mode = constants.DEBUG
	case constants.DEVELOP:
		c.Server.Mode = constants.TEST
	case constants.PRODUCTION:
		c.Server.Mode = constants.RELEASE
	case constants.EMPTY:
		c.Server.Mode = constants.DEBUG
	default:
		err = errors.New("server mode not valid: " + c.BaseConfig.Environment)
	}

	return
}

func (c *Config) getRole() (err error) {
	role := os.Getenv(constants.ROLE)

	switch role {
	case constants.WRITE, constants.READ, constants.WORKER:
		c.BaseConfig.Role = role
	case constants.EMPTY:
		c.BaseConfig.Role = constants.READ
	default:
		err = errors.New("role not valid: " + role)
	}

	return
}
