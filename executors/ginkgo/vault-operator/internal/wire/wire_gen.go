// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/k8s"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/profiles"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/temporary"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/vault"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/vaultoperator"
	"net/http"
)

// Injectors from injectors.go:

func SetUp(contextContext context.Context) (*Test, error) {
	config, err := profiles.NewConfig()
	if err != nil {
		return nil, err
	}

	restConfig, err := k8s.NewConfig(config)
	if err != nil {
		return nil, err
	}

	clientset, err := k8s.NewClientset(restConfig)
	if err != nil {
		return nil, err
	}

	coreV1Interface := k8s.NewCoreV1Client(clientset)
	namespace := temporary.NewNamespace(coreV1Interface)

	reader, err := k8s.NewManifestReader(restConfig)
	if err != nil {
		return nil, err
	}

	client := vaultoperator.NewClient(coreV1Interface, reader)
	authenticator := vault.NewAuthenticator(config)

	vaultConfig, err := vault.NewConfig()
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{}

	vaultClient, err := vault.NewClient(contextContext, vaultConfig, httpClient)
	if err != nil {
		return nil, err
	}

	secrets := vault.NewSecrets(authenticator, vaultClient)
	test := &Test{
		TemporaryNamespace: namespace,
		VaultOperator:      client,
		VaultSecrets:       secrets,
	}

	return test, nil
}
