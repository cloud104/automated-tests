package config

import (
	corev1 "k8s.io/api/core/v1"
)

type Crossplane struct {
	KubernetesProviderName string
	Namespace              string
}

func NewCrossplane(namespace *corev1.Namespace) *Crossplane {
	return &Crossplane{
		KubernetesProviderName: "kubernetes-provider",
		Namespace:              namespace.GetName(),
	}
}
