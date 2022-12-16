package config

import (
	"github.com/caarlos0/env/v6"
)

type AplicationConfig struct {
	DBConfig  DBConfig
	AppConfig AppConfig
}

type DBConfig struct {
	User     string `env:"DB_USER" envDefault:"root"`
	Passwd   string `env:"DB_PASS" envDefault:"*P*ndor*2018*"`
	ConnType string `env:"CONN_TYPE" envDefault:"tcp"`
	HostName string `env:"HOST_NAME" envDefault:"localhost:3306"`
	DBName   string `env:"DB_NAME" envDefault:"rampup"`
}

type AppConfig struct {
	GrpcAddr string `env:"GRPCADDR" envDefault:"localhost:9090"`
	HttpAddr string `env:"HTTPADDR" envDefault:"localhost:8080"`
}

func New() (AplicationConfig, error) {
	dbconfig := DBConfig{}
	if err := env.Parse(&dbconfig); err != nil {
		return AplicationConfig{}, err
	}
	appconfig := AppConfig{}
	if err := env.Parse(&appconfig); err != nil {
		return AplicationConfig{}, err
	}
	return AplicationConfig{DBConfig: dbconfig, AppConfig: appconfig}, nil
}
