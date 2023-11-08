package k8s

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewConfig() (*rest.Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to determine user's home directory: %w", err)
	}

	kubeconfig := filepath.Join(homeDir, ".kube", "config")

	rc, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client configuration: %w", err)
	}

	return rc, nil
}
