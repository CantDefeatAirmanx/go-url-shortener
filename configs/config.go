package configs

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var config Config

func LoadConfig() *Config {
	fmt.Println(
		os.Getenv("FOO_0_STR"),
	)

	var environment Environment = Environment(os.Getenv("GO_ENV"))
	if environment == "" {
		environment = DEV

		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	} else {
		environment = PROD
	}

	if err := env.Parse(&config); err != nil {
		panic(err)
	}

	return &config
}

func GetConfig() *Config {
	return &config
}
