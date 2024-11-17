package configs

type Config struct {
	GO_ENV   Environment `env:"GO_ENV"`
	APP_PORT string      `env:"APP_PORT"`
	Db       DbConfig    `envPrefix:"db__"`
	Auth     AuthConfig  `envPrefix:"auth__"`
}

type Test struct {
	Str string `env:"STR"`
	Num int    `env:"NUM"`
}

type Environment string

const (
	DEV  Environment = "development"
	PROD Environment = "production"
)

type DbConfig struct {
	Dsn string `env:"dsn"`
}

type AuthConfig struct {
	Secret string `env:"secret"`
}
