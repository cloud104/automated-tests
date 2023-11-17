package crossplane

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"github.com/totvs-cloud/go-manifest"

	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/config"
)

//go:embed manifests.yml
var manifestBytes []byte

type Client struct {
	config         *config.Crossplane
	manifestReader *manifest.Reader
}

func NewClient(config *config.Crossplane, manifestReader *manifest.Reader) *Client {
	return &Client{config: config, manifestReader: manifestReader}
}

func (c *Client) ApplyProviderManifests(ctx context.Context) (manifest.List, error) {
	m, err := c.manifestReader.FromBytes(manifestBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	m = m.Filter(ByAPIVersion("pkg.crossplane.io/v1"), ByKind("Provider"))

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
	if len(currentRevision) == 0 {
		return nil, errors.New("currentRevision cannot be empty")
	}

	m, err := c.manifestReader.FromBytes(manifestBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	m = m.Filter(ByAPIVersion("rbac.authorization.k8s.io/v1"))

	m, err = m.Transform(NamespaceTo(c.config.Namespace))
	if err != nil {
		return nil, err
	}

	m, err = m.Transform(RoleBindingSubjectTo(currentRevision))
	if err != nil {
		return nil, err
	}

	if err = m.Apply(ctx); err != nil {
		return nil, fmt.Errorf("failed to apply Provider manifest: %w", err)
	}

	return m, nil
}

func (c *Client) ApplyProviderConfigManifests(ctx context.Context) (manifest.List, error) {
	m, err := c.manifestReader.FromBytes(manifestBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}

	m = m.Filter(ByAPIVersion("kubernetes.crossplane.io/v1alpha1"), ByKind("ProviderConfig"))

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
	// TODO implement me
	panic("implement me")
}
