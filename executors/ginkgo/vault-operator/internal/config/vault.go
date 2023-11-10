package config

import (
	"fmt"
	"net/url"

	gofakeit "github.com/brianvoe/gofakeit/v6"
	"github.com/caarlos0/env/v9"
	corev1 "k8s.io/api/core/v1"
)

type Vault struct {
	Address   *url.URL
	Username  string
	Password  string
	Namespace string
}

func NewVault(namespace *corev1.Namespace) (*Vault, error) {
	cfg := &vaultFromEnv{}
	if err := env.ParseWithOptions(cfg, env.Options{Prefix: "VAULT_"}); err != nil {
		return nil, fmt.Errorf("error parsing Vault environment variables: %w", err)
	}

	addr, err := cfg.getAddressOrDefault(namespace.GetName())
	if err != nil {
		return nil, err
	}

	return &Vault{
		Address:   addr,
		Username:  cfg.getUsernameOrDefault(),
		Password:  cfg.getPasswordOrDefault(),
		Namespace: namespace.GetName(),
	}, nil
}

type vaultFromEnv struct {
	Address  *url.URL `env:"ADDRESS"`
	Username *string  `env:"USERNAME"`
	Password *string  `env:"PASSWORD"`
}

func (v *vaultFromEnv) getAddressOrDefault(namespace string) (*url.URL, error) {
	if v.Address == nil {
		def, err := url.Parse(fmt.Sprintf("http://vault-test.%s:8200", namespace))
		if err != nil {
			return nil, fmt.Errorf(": %w", err)
		}

		return def, nil
	}

	return v.Address, nil
}

func (v *vaultFromEnv) getUsernameOrDefault() string {
	if v.Username != nil {
		return *v.Username
	}

	return gofakeit.Username()
}

func (v *vaultFromEnv) getPasswordOrDefault() string {
	if v.Password != nil {
		return *v.Password
	}

	return gofakeit.Password(true, true, true, false, false, 8)
}
