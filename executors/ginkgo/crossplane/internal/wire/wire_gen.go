// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/crossplane"
	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/k8s"
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

	client, err := k8s.NewHTTPClient(restConfig)
	if err != nil {
		return nil, nil, err
	}

	clientset, err := k8s.NewClientset(restConfig, client)
	if err != nil {
		return nil, nil, err
	}

	coreV1Interface := k8s.NewCoreV1Client(clientset)

	namespace, cleanup, err := k8s.NewNamespace(ctx, basename, coreV1Interface)
	if err != nil {
		return nil, nil, err
	}

	configCrossplane := config.NewCrossplane(namespace)

	dynamicClient, err := k8s.NewDynamicClient(restConfig, client)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	reader, err := k8s.NewManifestReader(restConfig, client)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	crossplaneClient := crossplane.NewClient(configCrossplane, coreV1Interface, dynamicClient, reader)

	test, err := config.NewTest()
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	wireTest := &Test{
		Crossplane: crossplaneClient,
		Config:     test,
	}

	return wireTest, func() {
		cleanup()
	}, nil
}
