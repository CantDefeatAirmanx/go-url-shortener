package configs

type Config struct {
	GO_ENV   Environment
	APP_PORT string
	Db       DbConfig
}
type DbConfig struct {
	Dsn string
}

type Environment string

const (
	DEV  Environment = "development"
	PROD Environment = "production"
)
