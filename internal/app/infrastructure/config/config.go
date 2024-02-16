package config

import (
	_ "embed"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var settingsFile []byte

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Config struct {
	Port string   `yaml:"port"`
	DB   DBConfig `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	var config Config

	err := yaml.Unmarshal(settingsFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
