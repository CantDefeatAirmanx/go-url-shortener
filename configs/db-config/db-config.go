package dbconfig

type DbConfig struct {
	Dsn string `env:"dsn"`
}
