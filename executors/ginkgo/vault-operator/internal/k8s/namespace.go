package k8s

import (
	"context"
	"fmt"
	"log/slog"

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

func NewNamespace(ctx context.Context, basename string, k8sClient typedcorev1.CoreV1Interface) (*corev1.Namespace, func(), error) {
	name := generateName(basename)

	namespace, err := k8sClient.Namespaces().Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{
		Name: name,
		Labels: map[string]string{
			"app.kubernetes.io/name":       "vault-operator-test",
			"app.kubernetes.io/managed-by": "tks",
		},
	}}, metav1.CreateOptions{})
	if err != nil {
		return nil, func() {}, fmt.Errorf("failed to create Kubernetes namespace: %w", err)
	}

	slog.Info("Created Kubernetes namespace successfully.", "namespace", namespace.Name)

	cleanup := func() {
		if err = k8sClient.Namespaces().Delete(context.Background(), name, metav1.DeleteOptions{}); err != nil {
			slog.Warn("Failed to delete Kubernetes namespace. Manual cleanup may be required.", "namespace", name)
		}

		slog.Info("Deleted Kubernetes namespace successfully.", "namespace", name)
	}

	return namespace, cleanup, nil
}

func generateName(base string) string {
	if len(base) > maxGeneratedNameLength {
		base = base[:maxGeneratedNameLength]
	}

	return fmt.Sprintf("%s-%s", base, utilrand.String(randomLength))
}
