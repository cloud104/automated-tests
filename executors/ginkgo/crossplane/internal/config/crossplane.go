package config

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	utilrand "k8s.io/apimachinery/pkg/util/rand"
)

const (
	maxNameLength          = 63
	randomLength           = 5
	maxGeneratedNameLength = maxNameLength - randomLength
)

type Crossplane struct {
	KubernetesProviderName string
	Namespace              string
}

func NewCrossplane(namespace *corev1.Namespace) *Crossplane {
	return &Crossplane{
		KubernetesProviderName: generateName("kubernetes-provider"),
		Namespace:              namespace.GetName(),
	}
}

func generateName(base string) string {
	if len(base) > maxGeneratedNameLength {
		base = base[:maxGeneratedNameLength]
	}

	return fmt.Sprintf("%s-%s", base, utilrand.String(randomLength))
}
