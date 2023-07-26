package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Roles map[string]Role `yaml:"roles"`
}

type Role struct {
	Entitlements []string `yaml:"entitlements"`
}

func Load(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
