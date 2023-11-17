package wire

import (
	"github.com/google/wire"

	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/crossplane"
	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/k8s"
)

//nolint:unused // This function is used during compile-time to generate code for dependency injection
var providers = wire.NewSet(
	config.NewCrossplane,
	config.NewProfile,
	config.NewTest,
	crossplane.NewClient,
	k8s.NewClientset,
	k8s.NewConfig,
	k8s.NewCoreV1Client,
	k8s.NewManifestReader,
	k8s.NewNamespace,
	wire.Struct(new(Test), "*"),
)

type Test struct {
	Crossplane *crossplane.Client
	Config     *config.Test
}
