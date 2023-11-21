package k8s

import (
	"net/http"

	"k8s.io/client-go/rest"
)

func NewHTTPClient(config *rest.Config) (*http.Client, error) {
	client, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}
