package vault_operator_test

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/lmittmann/tint"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/k8s"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/temporary"
	intlvault "github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/vault"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/vaultoperator"
)

func TestVaultOperator(t *testing.T) {
	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  true,
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	}))
	slog.SetDefault(logger)

	RegisterFailHandler(Fail)
	RunSpecs(t, "VaultOperator Suite")
}

type Testing struct {
	TemporaryNamespace *temporary.Namespace
	VaultOperator      *vaultoperator.Client
	VaultSecrets       vault.Secrets
}

func initialize(contextContext context.Context) (*Testing, error) {
	config, err := k8s.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes configuration: %w", err)
	}

	clientset, err := k8s.NewClientset(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes clientset: %w", err)
	}

	coreClient := k8s.NewCoreV1Client(clientset)
	temporaryNamespace := temporary.NewNamespace(coreClient)

	manifestReader, err := k8s.NewManifestReader(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create manifest reader: %w", err)
	}

	vaultOperatorClient := vaultoperator.NewClient(coreClient, manifestReader)

	vaultConfig, err := intlvault.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create Vault configuration: %w", err)
	}

	httpClient := &http.Client{}

	vaultClient, err := intlvault.NewClient(contextContext, vaultConfig, httpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create Vault client: %w", err)
	}

	vaultSecrets := intlvault.NewSecrets(vaultClient)

	return &Testing{
		TemporaryNamespace: temporaryNamespace,
		VaultOperator:      vaultOperatorClient,
		VaultSecrets:       vaultSecrets,
	}, nil
}
