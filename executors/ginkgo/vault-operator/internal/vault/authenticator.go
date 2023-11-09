package vault

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"

	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/profiles"
)

const serviceAccountTokenPath = "/var/run/secrets/kubernetes.io/serviceaccount/token"

type Authenticator interface {
	Authenticate(ctx context.Context, client *vault.Client) error
}

func NewAuthenticator(config *profiles.Config) Authenticator {
	if strings.EqualFold(config.Active, "local") {
		return new(userpassAuthenticator)
	}

	return new(kubernetesAuthenticator)
}

type kubernetesAuthenticator struct{}

func (k *kubernetesAuthenticator) Authenticate(ctx context.Context, client *vault.Client) error {
	jwt, err := os.ReadFile(serviceAccountTokenPath)
	if err != nil {
		return fmt.Errorf("unable to read file containing service account token: %w", err)
	}

	resp, err := client.Auth.KubernetesLogin(ctx, schema.KubernetesLoginRequest{Role: "default", Jwt: string(jwt)})
	if err != nil {
		return fmt.Errorf("failed to authenticate with Vault: %w", err)
	}

	if err = client.SetToken(resp.Auth.ClientToken); err != nil {
		return fmt.Errorf("failed to set Vault client token: %w", err)
	}

	return nil
}

type userpassAuthenticator struct{}

func (u *userpassAuthenticator) Authenticate(ctx context.Context, client *vault.Client) error {
	// TODO: make this credentials dynamic if not set in env vars
	resp, err := client.Auth.UserpassLogin(ctx, "admin", schema.UserpassLoginRequest{Password: "admin"})
	if err != nil {
		return fmt.Errorf("failed to authenticate with Vault: %w", err)
	}

	if err = client.SetToken(resp.Auth.ClientToken); err != nil {
		return fmt.Errorf("failed to set Vault client token: %w", err)
	}

	return nil
}
