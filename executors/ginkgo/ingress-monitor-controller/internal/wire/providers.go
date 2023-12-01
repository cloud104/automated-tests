package wire

import (
	"net/http"

	"github.com/google/wire"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"

	"github.com/cloud104/automated-tests/executors/ginkgo/ingress-monitor-controller/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/ingress-monitor-controller/internal/k8s"
)

//nolint:unused // This function is used during compile-time to generate code for dependency injection
var providers = wire.NewSet(
	config.NewProfile,
	config.NewTest,
	k8s.NewAppsV1Client,
	k8s.NewClientset,
	k8s.NewConfig,
	wire.Struct(new(http.Client)),
	wire.Struct(new(Test), "*"),
)

type Test struct {
	Config *config.Test
	K8s    typedappsv1.AppsV1Interface
}
