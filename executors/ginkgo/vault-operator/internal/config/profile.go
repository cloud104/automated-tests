package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
)

type Profile struct {
	Active string `env:"ACTIVE" envDefault:"kubernetes"`
}

func NewProfile() (*Profile, error) {
	cfg := &Profile{}
	if err := env.ParseWithOptions(cfg, env.Options{Prefix: "PROFILE_"}); err != nil {
		return nil, fmt.Errorf("error parsing profile environment variables: %w", err)
	}

	return cfg, nil
}
