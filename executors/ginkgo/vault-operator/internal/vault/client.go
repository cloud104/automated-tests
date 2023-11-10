package vault

import (
	"context"
	"fmt"
	"net/http"

	vault "github.com/hashicorp/vault-client-go"

	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/config"
)

func NewClient(ctx context.Context, config *config.Vault, httpClient *http.Client) (*vault.Client, error) {
	client, err := vault.New(
		vault.WithAddress(config.Address.String()),
		vault.WithHTTPClient(httpClient),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Vault client: %w", err)
	}

	return client, nil
}
