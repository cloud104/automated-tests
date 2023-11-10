package wire

import (
	"net/http"

	"github.com/google/wire"

	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/k8s"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/vault"
	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/vaultoperator"
)

//nolint:unused // This function is used during compile-time to generate code for dependency injection
var providers = wire.NewSet(
	config.NewProfile,
	config.NewTest,
	config.NewVault,
	k8s.NewClientset,
	k8s.NewConfig,
	k8s.NewCoreV1Client,
	k8s.NewManifestReader,
	k8s.NewNamespace,
	vault.NewClient,
	vault.NewSecrets,
	vaultoperator.NewClient,
	wire.Struct(new(http.Client)),
	wire.Struct(new(Test), "*"),
)

type Test struct {
	Secrets  vault.Secrets
	Operator *vaultoperator.Client
	Config   *config.Test
}
