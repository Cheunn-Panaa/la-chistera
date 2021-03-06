package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config defines the application configuration each service conf is defined as another struct
type Config struct {
	Db  DatabaseConf `yaml:"database"`
	Env string       `yaml:"env"`
}

type DatabaseConf struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// NewConfig is constructor
func NewConfig(filename string) (config *Config, err error) {
	b, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(b, &config)
	return
}
