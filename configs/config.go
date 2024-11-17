package configs

import (
	"os"

	"github.com/joho/godotenv"
)

var config Config

func LoadConfig() *Config {
	var environment Environment = Environment(os.Getenv("GO_ENV"))
	if environment == "" {
		environment = DEV

		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	result := Config{
		GO_ENV: environment,
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		APP_PORT: os.Getenv("APP_PORT"),
	}
	config = result

	return &result
}

func GetConfig() *Config {
	return &config
}
