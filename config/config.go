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
	Version string `json:"version"`
}

type Postgres struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func ParseConfig(path string) (cfg *Config, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, cfg)
	return
}
