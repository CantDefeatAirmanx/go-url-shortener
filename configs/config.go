package configs

import (
	"fmt"
	"os"
	authconfig "url_shortener/configs/auth-config"
	dbconfig "url_shortener/configs/db-config"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var config Config

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

func LoadConfig() (*Config, error) {
	fmt.Println(
		os.Getenv("FOO_0_STR"),
	)

	var environment Environment = Environment(os.Getenv("GO_ENV"))
	if environment == "" || environment == DEV {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func GetConfig() *Config {
	return &config
}
