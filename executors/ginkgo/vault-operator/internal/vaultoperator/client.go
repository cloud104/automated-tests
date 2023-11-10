package vaultoperator

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/buger/jsonparser"
	"github.com/totvs-cloud/go-manifest"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/collections"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/k8s"
)

//go:embed manifests.yml
var manifestBytes []byte

type Client struct {
	config         *config.Vault
	k8sClient      typedcorev1.CoreV1Interface
	manifestReader *manifest.Reader
}

func NewClient(config *config.Vault, k8sClient typedcorev1.CoreV1Interface, manifestReader *manifest.Reader) *Client {
	return &Client{config: config, k8sClient: k8sClient, manifestReader: manifestReader}
}

func (c *Client) ApplyManifests(ctx context.Context) (manifest.List, error) {
	m, err := c.manifestReader.FromBytes(manifestBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	m, err = m.Transform(namespaceTo(c.config.Namespace))
	if err != nil {
		return nil, fmt.Errorf("failed to transform manifest: %w", err)
	}

	m, err = m.Transform(credentialsTo(c.config.Username, c.config.Password))
	if err != nil {
		return nil, fmt.Errorf("failed to transform manifest: %w", err)
	}

	if err = m.Apply(ctx); err != nil {
		return nil, fmt.Errorf("failed to apply manifest: %w", err)
	}

	return m, nil
}

func (c *Client) CountReadyPods(ctx context.Context) (int, error) {
	labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
		"app.kubernetes.io/name":             "vault",
		"statefulset.kubernetes.io/pod-name": "vault-test-0",
		"vault_cr":                           "vault-test",
	}}

	list, err := c.k8sClient.Pods(c.config.Namespace).List(ctx, metav1.ListOptions{
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

func credentialsTo(username, password string) func(u *unstructured.Unstructured) error {
	return func(u *unstructured.Unstructured) error {
		if u.GetAPIVersion() != "vault.banzaicloud.com/v1alpha1" || u.GetKind() != "Vault" {
			return nil
		}

		users := Users{{
			Username:      username,
			Password:      password,
			TokenPolicies: "allow_secrets",
		}}

		uj, err := users.Marshal()
		if err != nil {
			return err
		}

		data, err := u.MarshalJSON()
		if err != nil {
			return err
		}

		if data, err = jsonparser.Set(data, uj, "spec", "externalConfig", "auth", "[0]", "users"); err != nil {
			return err
		}

		return u.UnmarshalJSON(data)
	}
}
