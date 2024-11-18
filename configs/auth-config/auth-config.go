package authconfig

type AuthConfig struct {
	Secret string `env:"secret"`
}
