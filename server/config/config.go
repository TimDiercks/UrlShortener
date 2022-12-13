package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Environment string `json:"environment"`
	DB          string `json:"db"`
	Bind        string `json:"bind"`
}

func Parse(file string) (*Config, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(buf, &config); err != nil {
		return nil, err
	}

	return &config, err
}
