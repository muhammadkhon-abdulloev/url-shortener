package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	App      App      `json:"app"`
	Postgres Postgres `json:"postgres"`
	Server   Server   `json:"server"`
}

type App struct {
	BaseURL string `json:"baseURL" validate:"required"`
	IsProd  bool   `json:"isProd" validate:"required"`
	Version string `json:"version"`
}

type Postgres struct {
	Host     string `json:"host" validate:"required"`
	Port     string `json:"port" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	DBName   string `json:"dbName" validate:"required"`
}

type Server struct {
	Host string `json:"host" validate:"required"`
	Port string `json:"port" validate:"required"`
}

func ParseConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
