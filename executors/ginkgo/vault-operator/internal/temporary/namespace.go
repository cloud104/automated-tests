package temporary

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilrand "k8s.io/apimachinery/pkg/util/rand"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

const (
	maxNameLength          = 63
	randomLength           = 5
	maxGeneratedNameLength = maxNameLength - randomLength
)

type Namespace struct {
	k8sClient typedcorev1.CoreV1Interface
}

func NewNamespace(k8sClient typedcorev1.CoreV1Interface) *Namespace {
	return &Namespace{k8sClient: k8sClient}
}

func (n *Namespace) Create(ctx context.Context, base string) (string, error) {
	name := n.generateName(base)

	namespace := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{
		Name: name,
		Labels: map[string]string{
			"app.kubernetes.io/name":       "vault-operator-test",
			"app.kubernetes.io/managed-by": "tks",
		},
	}}
	if _, err := n.k8sClient.Namespaces().Create(ctx, namespace, metav1.CreateOptions{}); err != nil {
		return "", fmt.Errorf("failed to create Kubernetes namespace: %w", err)
	}

	return name, nil
}

func (n *Namespace) CleanUp(ctx context.Context, name string) error {
	if err := n.k8sClient.Namespaces().Delete(ctx, name, metav1.DeleteOptions{}); err != nil {
		return fmt.Errorf("failed to delete Kubernetes namespace %q: %w", name, err)
	}

	return nil
}

func (n *Namespace) generateName(base string) string {
	if len(base) > maxGeneratedNameLength {
		base = base[:maxGeneratedNameLength]
	}

	return fmt.Sprintf("%s-%s", base, utilrand.String(randomLength))
}
