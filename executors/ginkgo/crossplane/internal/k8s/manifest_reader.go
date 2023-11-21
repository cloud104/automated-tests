package k8s

import (
	"fmt"
	"net/http"

	"github.com/totvs-cloud/go-manifest"
	"k8s.io/client-go/rest"
)

const fieldManager = "crossplane-test"

func NewManifestReader(config *rest.Config, httpClient *http.Client) (*manifest.Reader, error) {
	mr, err := manifest.NewReaderForConfigAndClient(fieldManager, config, httpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create a Kubernetes manifest reader: %w", err)
	}

	return mr, nil
}
