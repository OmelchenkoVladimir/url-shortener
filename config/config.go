package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Db struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"db"`
	Server struct {
		Port string `json:"port"`
		Url  string `json:"url"`
	} `json:"server"`
}

func ReadConfig(path string) (*Config, error) {
	config, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var res Config
	if err := json.Unmarshal(config, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
