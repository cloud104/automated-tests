package k8s

import (
	"net/http"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

func NewDynamicClient(config *rest.Config, httpClient *http.Client) (*dynamic.DynamicClient, error) {
	dynamicClient, err := dynamic.NewForConfigAndClient(config, httpClient)
	if err != nil {
		return nil, err
	}

	return dynamicClient, nil
}
