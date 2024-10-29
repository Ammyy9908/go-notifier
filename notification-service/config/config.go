package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const DevEnvironment = "dev"

type Config struct {
	Map    map[string]interface{}
	Errors []error
}

var (
	EmptyPropertySource = errors.New("no configurations in property source")
	SourceHasNoConfigs  = errors.New("source has no configurations inside")
)

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

// AddEnv adds a given environment key value pairs into the Config structs map, overriding existing key values
func (c *Config) AddEnv(configurationsMap map[string]interface{}) {
	for k, v := range configurationsMap {
		if _, ok := v.(string); ok {
			value := v.(string)
			if len(value) >= 2 && value[:2] == "${" {
				c.Map[k] = os.Getenv(value[2 : len(value)-1])
				continue
			}
		}
	}
}

// HasErrors returns true if the creation of configuration resulted in at least one error
func (c *Config) HasErrors() bool {
	return len(c.Errors) > 0
}

// SetViper sets Config map as Viper default variables
func (c *Config) SetViper() {
	for k, v := range c.Map {
		viper.SetDefault(k, v)
	}
}

// FromConfigFile returns a Config with a populated map of values from a local SpringCloud configuration file
func (c *Config) FromConfigFile(path string) *Config {
	springConfig, err := GetFile(path)
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

// GetFile fetches the byte slice from a specified file
func GetFile(path string) (SpringResponse, error) {
	var cloudConfig SpringResponse
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Not able to read file from path", err)
		return SpringResponse{}, err
	}
	err = json.Unmarshal(file, &cloudConfig)
	if err != nil {
		return SpringResponse{}, err
	}
	return cloudConfig, nil
}

// SpringResponse Structs having same structure as want from Cloud config
type SpringResponse struct {
	Name            string           `json:"name"`
	Profiles        []string         `json:"profiles"`
	Label           string           `json:"label"`
	Version         string           `json:"version"`
	PropertySources []propertySource `json:"propertySources"`
}

type propertySource struct {
	Name   string                 `json:"name"`
	Source map[string]interface{} `json:"source"`
}

// ToMap parses the springResponse property sources struct into a map
func (cc *SpringResponse) ToMap() (map[string]interface{}, error) {
	if len(cc.PropertySources) < 1 {
		return nil, EmptyPropertySource
	}
	if cc.PropertySources[0].Source == nil {
		return nil, SourceHasNoConfigs
	}
	return cc.PropertySources[0].Source, nil
}
