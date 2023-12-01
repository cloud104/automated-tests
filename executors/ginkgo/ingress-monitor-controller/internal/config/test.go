package config

import (
	"fmt"
	"time"

	env "github.com/caarlos0/env/v10"
)

type Test struct {
	Timeout   time.Duration `env:"TEST_TIMEOUT" envDefault:"1m"`
	Namespace string        `env:"NAMESPACE" envDefault:"tks-system"`
}

func NewTest() (*Test, error) {
	cfg := &Test{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("error parsing Vault environment variables: %w", err)
	}

	return cfg, nil
}
