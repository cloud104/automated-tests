package vault

import (
	"context"
	"sync"

	vault "github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

type Secrets interface {
	KvV2Configure(ctx context.Context, request schema.KvV2ConfigureRequest, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error)
	KvV2Delete(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error)
	KvV2DeleteMetadataAndAllVersions(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error)
	KvV2DeleteVersions(ctx context.Context, path string, request schema.KvV2DeleteVersionsRequest, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error)
	KvV2DestroyVersions(ctx context.Context, path string, request schema.KvV2DestroyVersionsRequest, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error)
	KvV2List(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[schema.StandardListResponse], error)
	KvV2Read(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[schema.KvV2ReadResponse], error)
	KvV2ReadConfiguration(ctx context.Context, options ...vault.RequestOption) (*vault.Response[schema.KvV2ReadConfigurationResponse], error)
	KvV2ReadMetadata(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[schema.KvV2ReadMetadataResponse], error)
	KvV2ReadSubkeys(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[schema.KvV2ReadSubkeysResponse], error)
	KvV2UndeleteVersions(ctx context.Context, path string, request schema.KvV2UndeleteVersionsRequest, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error)
	KvV2Write(ctx context.Context, path string, request schema.KvV2WriteRequest, options ...vault.RequestOption) (*vault.Response[schema.KvV2WriteResponse], error)
	KvV2WriteMetadata(ctx context.Context, path string, request schema.KvV2WriteMetadataRequest, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error)
}

func NewSecrets(auth Authenticator, client *vault.Client) Secrets {
	return &secretsAuthenticatedOnce{auth: auth, client: client}
}

type secretsAuthenticatedOnce struct {
	once   sync.Once
	auth   Authenticator
	client *vault.Client
}

func (s *secretsAuthenticatedOnce) authenticateOnce(ctx context.Context) error {
	var err error

	s.once.Do(func() {
		err = s.auth.Authenticate(ctx, s.client)
	})

	return err
}

func (s *secretsAuthenticatedOnce) KvV2Configure(ctx context.Context, request schema.KvV2ConfigureRequest, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2Configure(ctx, request, options...)
}

func (s *secretsAuthenticatedOnce) KvV2Delete(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2Delete(ctx, path, options...)
}

func (s *secretsAuthenticatedOnce) KvV2DeleteMetadataAndAllVersions(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2Delete(ctx, path, options...)
}

func (s *secretsAuthenticatedOnce) KvV2DeleteVersions(ctx context.Context, path string, request schema.KvV2DeleteVersionsRequest, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2DeleteVersions(ctx, path, request, options...)
}

func (s *secretsAuthenticatedOnce) KvV2DestroyVersions(ctx context.Context, path string, request schema.KvV2DestroyVersionsRequest, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2DestroyVersions(ctx, path, request, options...)
}

func (s *secretsAuthenticatedOnce) KvV2List(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[schema.StandardListResponse], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2List(ctx, path, options...)
}

func (s *secretsAuthenticatedOnce) KvV2Read(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[schema.KvV2ReadResponse], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2Read(ctx, path, options...)
}

func (s *secretsAuthenticatedOnce) KvV2ReadConfiguration(ctx context.Context, options ...vault.RequestOption) (*vault.Response[schema.KvV2ReadConfigurationResponse], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2ReadConfiguration(ctx, options...)
}

func (s *secretsAuthenticatedOnce) KvV2ReadMetadata(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[schema.KvV2ReadMetadataResponse], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2ReadMetadata(ctx, path, options...)
}

func (s *secretsAuthenticatedOnce) KvV2ReadSubkeys(ctx context.Context, path string, options ...vault.RequestOption) (*vault.Response[schema.KvV2ReadSubkeysResponse], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2ReadSubkeys(ctx, path, options...)
}

func (s *secretsAuthenticatedOnce) KvV2UndeleteVersions(ctx context.Context, path string, request schema.KvV2UndeleteVersionsRequest, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2UndeleteVersions(ctx, path, request, options...)
}

func (s *secretsAuthenticatedOnce) KvV2Write(ctx context.Context, path string, request schema.KvV2WriteRequest, options ...vault.RequestOption) (*vault.Response[schema.KvV2WriteResponse], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2Write(ctx, path, request, options...)
}

func (s *secretsAuthenticatedOnce) KvV2WriteMetadata(ctx context.Context, path string, request schema.KvV2WriteMetadataRequest, options ...vault.RequestOption) (*vault.Response[map[string]interface{}], error) {
	if err := s.authenticateOnce(ctx); err != nil {
		return nil, err
	}

	return s.client.Secrets.KvV2WriteMetadata(ctx, path, request, options...)
}
