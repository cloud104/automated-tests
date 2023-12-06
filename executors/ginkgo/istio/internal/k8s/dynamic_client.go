package k8s

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

func NewDynamicClient(config *rest.Config) (*dynamic.DynamicClient, error) {
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return dynamicClient, nil
}
