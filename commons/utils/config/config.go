package config

import (
	ConfigCloud "go-notifier/commons/utils/config/internal/providers"
	"strings"

	"github.com/spf13/viper"
)

const DevEnvironment = "dev"

type Config struct {
	Map    map[string]interface{}
	Errors []error
}

// New creates a new Config struct
func New() *Config {
	var errors []error
	m := make(map[string]interface{})
	return &Config{Map: m, Errors: errors}
}

// AddMap adds a given map key value pairs into the Config structs map, overriding existing key values
func (c *Config) AddMap(configurationsMap map[string]interface{}) {
	for k, v := range configurationsMap {
		c.Map[k] = v
	}
}

// HasErrors returns true if the creation of configuration resulted in at least one error
func (c *Config) HasErrors() bool {
	if len(c.Errors) > 0 {
		return true
	}
	return false
}

// SetViper sets Config map as Viper default variables
func (c *Config) SetViper() {
	for k, v := range c.Map {
		viper.SetDefault(k, v)
	}
}

// FromS3Cloud builds a url from given parameters and returns the resulting config
func (c *Config) FromS3Cloud(service, env string, version string) *Config {
	if strings.EqualFold(env, DevEnvironment) {
		c.Errors = append(c.Errors, ConfigCloud.DevEnvironmentError)
		return c
	}
	fileName := env + "-" + service + ".json"
	url := ConfigCloud.BaseUrl + "/" + env + "/" + service + "/" + version + "/" + fileName
	return c.FromS3CloudUrl(url)

}

func (c *Config) FromS3CloudUrl(url string) *Config {
	springCloud, err := ConfigCloud.GetCloud(url)
	if err != nil {
		c.Errors = append(c.Errors, err)
		return c
	}
	springMap, err := springCloud.ToMap()
	if err != nil {
		c.Errors = append(c.Errors, err)
		return c
	}
	c.AddMap(springMap)
	return c
}

// FromFile returns a Config with a populated map of values from a local SpringCloud configuration file
func (c *Config) FromFile(path string) *Config {
	springConfig, err := ConfigCloud.GetFile(path)
	if err != nil {
		c.Errors = append(c.Errors, err)
		return c
	}
	springMap, err := springConfig.ToMap()
	if err != nil {
		c.Errors = append(c.Errors, err)
		return c
	}
	c.AddMap(springMap)
	return c
}
