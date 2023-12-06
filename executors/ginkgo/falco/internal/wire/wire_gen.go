// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/cloud104/automated-tests/executors/ginkgo/falco/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/falco/internal/k8s"
)

// Injectors from injectors.go:

func SetUp(ctx context.Context) (*Test, error) {
	profile, err := config.NewProfile()
	if err != nil {
		return nil, err
	}

	restConfig, err := k8s.NewConfig(profile)
	if err != nil {
		return nil, err
	}

	clientset, err := k8s.NewClientset(restConfig)
	if err != nil {
		return nil, err
	}

	appsV1Interface := k8s.NewAppsV1Client(clientset)
	coreV1Interface := k8s.NewCoreV1Client(clientset)

	test, err := config.NewTest()
	if err != nil {
		return nil, err
	}

	wireTest := &Test{
		AppsV1: appsV1Interface,
		CoreV1: coreV1Interface,
		Config: test,
	}

	return wireTest, nil
}
