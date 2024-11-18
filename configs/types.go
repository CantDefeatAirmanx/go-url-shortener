package configs

import (
	authconfig "url_shortener/configs/auth-config"
	dbconfig "url_shortener/configs/db-config"
)

type Config struct {
	GO_ENV   Environment           `env:"GO_ENV"`
	APP_PORT string                `env:"APP_PORT"`
	Db       dbconfig.DbConfig     `envPrefix:"db__"`
	Auth     authconfig.AuthConfig `envPrefix:"auth__"`
}

type Environment string

const (
	DEV  Environment = "development"
	PROD Environment = "production"
)
