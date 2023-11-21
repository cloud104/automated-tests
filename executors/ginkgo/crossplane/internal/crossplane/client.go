package crossplane

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log/slog"

	"github.com/buger/jsonparser"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/slogr"
	"github.com/totvs-cloud/go-manifest"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/collections"
	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/k8s"
)

//go:embed manifests.yml
var manifestBytes []byte

type Client struct {
	config         *config.Crossplane
	coreClient     typedcorev1.CoreV1Interface
	dynamicClient  dynamic.Interface
	manifestReader *manifest.Reader
}

func NewClient(config *config.Crossplane, coreClient typedcorev1.CoreV1Interface, dynamicClient dynamic.Interface, manifestReader *manifest.Reader) *Client {
	return &Client{config: config, coreClient: coreClient, dynamicClient: dynamicClient, manifestReader: manifestReader}
}

func (c *Client) ApplyProviderManifests(ctx context.Context) (manifest.List, error) {
	ctx = logr.NewContext(ctx, slogr.NewLogr(slog.Default().Handler()))

	m, err := c.manifestReader.FromBytes(manifestBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	m = m.Filter(manifest.All(
		manifest.ByAPIVersion("pkg.crossplane.io/v1"),
		manifest.ByKind("Provider"),
	))

	m, err = m.Transform(NameTo(c.config.KubernetesProviderName))
	if err != nil {
		return nil, err
	}

	if err = m.Apply(ctx); err != nil {
		return nil, fmt.Errorf("failed to apply Provider manifest: %w", err)
	}

	return m, nil
}

func (c *Client) ApplyRBACManifests(ctx context.Context, currentRevision string) (manifest.List, error) {
	ctx = logr.NewContext(ctx, slogr.NewLogr(slog.Default().Handler()))

	if len(currentRevision) == 0 {
		return nil, errors.New("currentRevision cannot be empty")
	}

	m, err := c.manifestReader.FromBytes(manifestBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	m = m.Filter(
		manifest.ByAPIVersion("rbac.authorization.k8s.io/v1"),
	)

	m, err = m.Transform(NamespaceTo(c.config.Namespace))
	if err != nil {
		return nil, err
	}

	serviceAccount, err := c.FindServiceAccount(ctx, currentRevision)
	if err != nil {
		return nil, err
	}

	m, err = m.Transform(RoleBindingSubjectTo(&Subject{
		Kind:      "ServiceAccount",
		Name:      serviceAccount.GetName(),
		Namespace: serviceAccount.GetNamespace(),
	}))
	if err != nil {
		return nil, err
	}

	if err = m.Apply(ctx); err != nil {
		return nil, fmt.Errorf("failed to apply Provider manifest: %w", err)
	}

	return m, nil
}

func (c *Client) ApplyProviderConfigManifests(ctx context.Context) (manifest.List, error) {
	ctx = logr.NewContext(ctx, slogr.NewLogr(slog.Default().Handler()))

	m, err := c.manifestReader.FromBytes(manifestBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	m = m.Filter(manifest.All(
		manifest.ByAPIVersion("kubernetes.crossplane.io/v1alpha1"),
		manifest.ByKind("ProviderConfig"),
	))

	m, err = m.Transform(NameTo(c.config.KubernetesProviderName))
	if err != nil {
		return nil, err
	}

	if err = m.Apply(ctx); err != nil {
		return nil, fmt.Errorf("failed to apply Provider manifest: %w", err)
	}

	return m, nil
}

func (c *Client) GetProviderStatus(ctx context.Context) (*ProviderStatus, error) {
	resource := schema.GroupVersionResource{Group: "pkg.crossplane.io", Version: "v1", Resource: "providers"}

	u, err := c.dynamicClient.Resource(resource).Get(ctx, c.config.KubernetesProviderName, metav1.GetOptions{})
	if k8serrors.IsNotFound(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	data, err := u.MarshalJSON()
	if err != nil {
		return nil, err
	}

	sj, _, _, err := jsonparser.Get(data, "status")
	if err != nil && "Key path not found" == err.Error() {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return UnmarshalProviderStatus(sj)
}

func (c *Client) ApplyObjectManifests(ctx context.Context) (manifest.List, error) {
	ctx = logr.NewContext(ctx, slogr.NewLogr(slog.Default().Handler()))

	m, err := c.manifestReader.FromBytes(manifestBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	m = m.Filter(manifest.All(
		manifest.ByAPIVersion("kubernetes.crossplane.io/v1alpha1"),
		manifest.ByKind("Object"),
	))

	m, err = m.Transform(NamespaceTo(c.config.Namespace))
	if err != nil {
		return nil, err
	}

	m, err = m.Transform(ObjectManifestNamespaceTo(c.config.Namespace))
	if err != nil {
		return nil, err
	}

	m, err = m.Transform(ObjectProviderConfigRefTo(c.config.KubernetesProviderName))
	if err != nil {
		return nil, err
	}

	m, err = m.Transform(ObjectProviderConfigRefTo(c.config.KubernetesProviderName))
	if err != nil {
		return nil, err
	}

	if err = m.Apply(ctx); err != nil {
		return nil, fmt.Errorf("failed to apply Provider manifest: %w", err)
	}

	return m, nil
}

func (c *Client) CountReadySamplePods(ctx context.Context) (int, error) {
	labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
		"app": "sample-pod",
	}}

	list, err := c.coreClient.Pods(c.config.Namespace).List(ctx, metav1.ListOptions{
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

func (c *Client) FindServiceAccount(ctx context.Context, revision string) (*corev1.ServiceAccount, error) {
	namespaces, err := c.coreClient.Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list namespaces: %w", err)
	}

	for _, namespace := range namespaces.Items {
		serviceAccounts, err := c.coreClient.ServiceAccounts(namespace.GetName()).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to list pods: %w", err)
		}

		for _, sa := range serviceAccounts.Items {
			owners := collections.Filter(
				sa.GetOwnerReferences(),
				ByProviderRevision(revision),
			)

			if len(owners) > 0 {
				return &sa, nil
			}
		}
	}

	return nil, nil
}

func (c *Client) DeleteManifests(ctx context.Context, manifests manifest.List) error {
	ctx = logr.NewContext(ctx, slogr.NewLogr(slog.Default().Handler()))

	objects := manifests.Filter(manifest.All(
		manifest.ByAPIVersion("kubernetes.crossplane.io/v1alpha1"),
		manifest.ByKind("Object"),
	))
	if err := objects.Delete(ctx, manifest.WaitForDelete()); err != nil {
		return err
	}

	otherK8s := manifests.Filter(manifest.All(
		manifest.Not(manifest.In(objects)),
		manifest.ByAPIVersion("kubernetes.crossplane.io/v1alpha1"),
	))
	if err := otherK8s.Delete(ctx, manifest.WaitForDelete()); err != nil {
		return err
	}

	remaining := manifests.Filter(
		manifest.Not(manifest.In(objects.Append(otherK8s))),
	)
	if err := remaining.Delete(ctx, manifest.WaitForDelete()); err != nil {
		return err
	}

	return nil
}

func ByProviderRevision(revision string) func(ownerReference metav1.OwnerReference, index int) bool {
	return func(ownerReference metav1.OwnerReference, index int) bool {
		return ownerReference.APIVersion == "pkg.crossplane.io/v1" &&
			ownerReference.Kind == "ProviderRevision" &&
			ownerReference.Name == revision
	}
}
