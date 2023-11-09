package k8s

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/profiles"
)

func NewConfig(config *profiles.Config) (*rest.Config, error) {
	if strings.EqualFold(config.Active, "local") {
		return localConfig()
	}

	return inClusterConfig()
}

func inClusterConfig() (*rest.Config, error) {
	rc, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve in-cluster Kubernetes client configuration: %w", err)
	}

	return rc, nil
}

func localConfig() (*rest.Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to determine user's home directory: %w", err)
	}

	kubeconfig := filepath.Join(homeDir, ".kube", "config")

	rc, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve local Kubernetes client configuration: %w", err)
	}

	return rc, nil
}
