package k8s

import (
	"fmt"

	"github.com/totvs-cloud/go-manifest"
	"k8s.io/client-go/rest"
)

const fieldManager = "crossplane-test"

func NewManifestReader(config *rest.Config) (*manifest.Reader, error) {
	mr, err := manifest.NewReader(fieldManager, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create a Kubernetes manifest reader: %w", err)
	}

	return mr, nil
}
