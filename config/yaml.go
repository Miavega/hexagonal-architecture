package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const configPath = "%s/config/%s.yml"

func (c *Config) getYaml() (err error) {
	basePath, err := os.Getwd()
	if err != nil {
		return
	}

	yamlPath := fmt.Sprintf(configPath, basePath, c.BaseConfig.Environment)
	file, err := os.Open(yamlPath)
	if err != nil {
		return
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err = d.Decode(c); err != nil {
		return
	}

	return
}
