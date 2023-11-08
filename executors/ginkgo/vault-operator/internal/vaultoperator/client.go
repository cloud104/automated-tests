package vaultoperator

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/totvs-cloud/go-manifest"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/collections"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/k8s"
)

//go:embed manifests.yml
var manifestBytes []byte

type Client struct {
	k8sClient      typedcorev1.CoreV1Interface
	manifestReader *manifest.Reader
}

func NewClient(k8sClient typedcorev1.CoreV1Interface, manifestReader *manifest.Reader) *Client {
	return &Client{k8sClient: k8sClient, manifestReader: manifestReader}
}

func (c *Client) ApplyManifests(ctx context.Context, namespace string) (manifest.List, error) {
	m, err := c.manifestReader.FromBytes(manifestBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	m, err = m.Transform(namespaceTo(namespace))
	if err != nil {
		return nil, fmt.Errorf("failed to transform manifest: %w", err)
	}

	if err = m.Apply(ctx); err != nil {
		return nil, fmt.Errorf("failed to apply manifest: %w", err)
	}

	return m, nil
}

func (c *Client) CleanUp(ctx context.Context, namespace string, manifests manifest.List) error {
	m, err := manifests.Transform(namespaceTo(namespace))
	if err != nil {
		return fmt.Errorf("failed to transform manifests: %w", err)
	}

	if err = m.Delete(ctx); err != nil {
		return fmt.Errorf("failed to delete resources: %w", err)
	}

	return nil
}

func (c *Client) CountReadyPods(ctx context.Context, namespace string) (int, error) {
	labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
		"app.kubernetes.io/name":             "vault",
		"statefulset.kubernetes.io/pod-name": "vault-test-0",
		"vault_cr":                           "vault-test",
	}}

	list, err := c.k8sClient.Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})
	if err != nil {
		return 0, fmt.Errorf("failed to list pods: %w", err)
	}

	pods := collections.Filter(list.Items, func(pod corev1.Pod, index int) bool {
		return k8s.IsPodStatusConditionTrue(pod.Status.Conditions, corev1.ContainersReady)
	})

	return len(pods), nil
}

func namespaceTo(namespace string) func(u *unstructured.Unstructured) error {
	return func(u *unstructured.Unstructured) error {
		if u.GetKind() != "ClusterRoleBinding" {
			u.SetNamespace(namespace)
		}

		return nil
	}
}
