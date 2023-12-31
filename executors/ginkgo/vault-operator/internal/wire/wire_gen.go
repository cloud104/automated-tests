// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/k8s"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/vault"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/vaultoperator"
	"net/http"
)

// Injectors from injectors.go:

func SetUp(ctx context.Context, basename string) (*Test, func(), error) {
	profile, err := config.NewProfile()
	if err != nil {
		return nil, nil, err
	}

	restConfig, err := k8s.NewConfig(profile)
	if err != nil {
		return nil, nil, err
	}

	clientset, err := k8s.NewClientset(restConfig)
	if err != nil {
		return nil, nil, err
	}

	coreV1Interface := k8s.NewCoreV1Client(clientset)

	namespace, cleanup, err := k8s.NewNamespace(ctx, basename, coreV1Interface)
	if err != nil {
		return nil, nil, err
	}

	configVault, err := config.NewVault(namespace)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	client := &http.Client{}

	vaultClient, err := vault.NewClient(ctx, configVault, client)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	secrets := vault.NewSecrets(vaultClient, configVault)

	reader, err := k8s.NewManifestReader(restConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	vaultoperatorClient := vaultoperator.NewClient(configVault, coreV1Interface, reader)

	test, err := config.NewTest()
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	wireTest := &Test{
		Secrets:  secrets,
		Operator: vaultoperatorClient,
		Config:   test,
	}

	return wireTest, func() {
		cleanup()
	}, nil
}
