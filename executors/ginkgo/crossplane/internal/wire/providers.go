package wire

import (
	"github.com/google/wire"
	corev1 "k8s.io/api/core/v1"

	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/k8s"
)

//nolint:unused // This function is used during compile-time to generate code for dependency injection
var providers = wire.NewSet(
	config.NewProfile,
	k8s.NewClientset,
	k8s.NewConfig,
	k8s.NewCoreV1Client,
	k8s.NewNamespace,
	wire.Struct(new(Test), "*"),
)

type Test struct {
	Namespace *corev1.Namespace
}
