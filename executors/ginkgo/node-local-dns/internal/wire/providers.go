package wire

import (
	"github.com/google/wire"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"github.com/cloud104/automated-tests/executors/ginkgo/node-local-dns/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/node-local-dns/internal/k8s"
)

//nolint:unused // This function is used during compile-time to generate code for dependency injection
var providers = wire.NewSet(
	config.NewProfile,
	config.NewTest,
	k8s.NewAppsV1Client,
	k8s.NewClientset,
	k8s.NewConfig,
	k8s.NewCoreV1Client,
	wire.Struct(new(Test), "*"),
)

type Test struct {
	AppsV1 appsv1.AppsV1Interface
	CoreV1 corev1.CoreV1Interface
	Config *config.Test
}
