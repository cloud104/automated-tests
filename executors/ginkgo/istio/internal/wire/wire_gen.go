// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"context"
	"github.com/cloud104/automated-tests/executors/ginkgo/istio/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/istio/internal/istiooperator"
	"github.com/cloud104/automated-tests/executors/ginkgo/istio/internal/k8s"
)

// Injectors from injectors.go:

func SetUp(ctx context.Context) (*Test, error) {
	test, err := config.NewTest()
	if err != nil {
		return nil, err
	}

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

	dynamicClient, err := k8s.NewDynamicClient(restConfig)
	if err != nil {
		return nil, err
	}

	client := istiooperator.NewClient(dynamicClient, test)
	wireTest := &Test{
		Config:        test,
		K8s:           appsV1Interface,
		IstioOperator: client,
	}

	return wireTest, nil
}