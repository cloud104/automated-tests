package vault

import (
	"fmt"
	"net/url"

	env "github.com/caarlos0/env/v9"
)

type Config struct {
	Address  url.URL `env:"ADDRESS" envDefault:"http://127.0.0.1:8200"`
	Username string  `env:"USERNAME" envDefault:"admin"`
	Password string  `env:"PASSWORD" envDefault:"admin"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.ParseWithOptions(cfg, env.Options{Prefix: "VAULT_"}); err != nil {
		return nil, fmt.Errorf("error parsing Vault environment variables: %w", err)
	}

	return cfg, nil
}
