package config

import (
	"sync"
	"time"
)

var (
	configuration Config
	once          sync.Once
)

// Config struct for webapp config
type Config struct {
	BaseConfig struct {
		Role        string `yaml:"role"`
		Environment string `yaml:"environment"`
		BasePath    string `yaml:"base_path"`
	} `yaml:"base_config"`

	Server struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`

		// Port is the local machine TCP Port to bind the HTTP Server to
		Port    string `yaml:"port"`
		Timeout struct {
			// Server is the general server timeout to use
			// for graceful shutdowns
			Server time.Duration `yaml:"server"`

			// Write is the amount of time to wait until an HTTP server
			// write opperation is cancelled
			Write time.Duration `yaml:"write"`

			// Read is the amount of time to wait until an HTTP server
			// read operation is cancelled
			Read time.Duration `yaml:"read"`

			// Read is the amount of time to wait
			// until an IDLE HTTP session is closed
			Idle time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
		// Server Mode (GIN: debug, release, test)
		Mode string `yaml:"mode"`
	} `yaml:"server"`
}

func Get() (*Config, error) {
	var err error
	once.Do(func() {
		err = configuration.BuildConfig()
	})
	return &configuration, err
}

func (c *Config) BuildConfig() (err error) {
	err = c.getEnv()
	err = c.getRole()
	err = c.getServerMode()
	err = c.getYaml()
	return
}
