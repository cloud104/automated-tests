package vault

import (
	"context"
	"fmt"
	"net/http"

	vault "github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func NewClient(ctx context.Context, config *Config, httpClient *http.Client) (*vault.Client, error) {
	client, err := vault.New(
		vault.WithAddress(config.Address.String()),
		vault.WithHTTPClient(httpClient),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Vault client: %w", err)
	}

	resp, err := client.Auth.UserpassLogin(ctx, config.Username, schema.UserpassLoginRequest{Password: config.Password})
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate with Vault: %w", err)
	}

	if err = client.SetToken(resp.Auth.ClientToken); err != nil {
		return nil, fmt.Errorf("failed to set Vault client token: %w", err)
	}

	return client, nil
}

func NewSecrets(client *vault.Client) vault.Secrets {
	return client.Secrets
}
