package k8s

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

func NewClientset(config *rest.Config) (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes clientset: %w", err)
	}

	return clientset, nil
}

func NewAppsV1Client(clientset *kubernetes.Clientset) appsv1.AppsV1Interface {
	client := clientset.AppsV1()
	return client
}

func NewCoreV1Client(clientset *kubernetes.Clientset) corev1.CoreV1Interface {
	client := clientset.CoreV1()
	return client
}
