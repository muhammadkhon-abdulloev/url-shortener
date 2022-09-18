package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Config struct {
	Logger      Logger      `json:"Logger"`
	Postgres    Postgres    `json:"Postgres"`
	PostgresDev PostgresDev `json:"PostgresDev"`
	Server      Server      `json:"Server"`
}

type Logger struct {
	Development       bool   `json:"Development"`
	DisableCaller     bool   `json:"DisableCaller"`
	DisableStacktrace bool   `json:"DisableStacktrace"`
	Encoding          string `json:"Encoding"`
	Level             string `json:"Level"`
}

type Postgres struct {
	Host              string `json:"Host" validate:"required"`
	Port              string `json:"Port" validate:"required"`
	Username          string `json:"Username" validate:"required"`
	Password          string `json:"Password" validate:"required"`
	DBName            string `json:"DBName" validate:"required"`
	PostgresqlSslmode bool   `json:"PostgresqlSslmode"`
	PgDriver          string `json:"PgDriver"`
}

type PostgresDev struct {
	Host              string `json:"Host" validate:"required"`
	Port              string `json:"Port" validate:"required"`
	Username          string `json:"Username" validate:"required"`
	Password          string `json:"Password" validate:"required"`
	DBName            string `json:"DBName" validate:"required"`
	PostgresqlSslmode bool   `json:"PostgresqlSslmode"`
	PgDriver          string `json:"PgDriver"`
}

type Server struct {
	AppVersion   string        `json:"AppVersion"`
	BaseURL      string        `json:"BaseURL" validate:"required"`
	DevPort      string        `json:"DevPort" validate:"required"`
	Mode         string        `json:"Mode" validate:"required"`
	Port         string        `json:"Port" validate:"required"`
	ReadTimeout  time.Duration `json:"ReadTimeout" validate:"required"`
	WriteTimeout time.Duration `json:"WriteTimeout" validate:"required"`
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
