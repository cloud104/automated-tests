package profiles

import (
	"fmt"

	env "github.com/caarlos0/env/v9"
)

type Config struct {
	Active string `env:"ACTIVE" envDefault:"kubernetes"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.ParseWithOptions(cfg, env.Options{Prefix: "PROFILES_"}); err != nil {
		return nil, fmt.Errorf("error parsing Vault environment variables: %w", err)
	}

	return cfg, nil
}
