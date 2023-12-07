package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v9"
)

type Test struct {
	SkipDelete bool          `env:"SKIP_DELETE" envDefault:"false"`
	Timeout    time.Duration `env:"TIMEOUT" envDefault:"1m"`
}

func NewTest() (*Test, error) {
	cfg := &Test{}
	if err := env.ParseWithOptions(cfg, env.Options{Prefix: "TEST_"}); err != nil {
		return nil, fmt.Errorf("error parsing test environment variables: %w", err)
	}

	return cfg, nil
}
