package wire

import (
	"net/http"

	"github.com/google/wire"
	"k8s.io/client-go/dynamic"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"

	"github.com/cloud104/automated-tests/executors/ginkgo/istio/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/istio/internal/istiooperator"
	"github.com/cloud104/automated-tests/executors/ginkgo/istio/internal/k8s"
)

//nolint:unused // This function is used during compile-time to generate code for dependency injection
var providers = wire.NewSet(
	config.NewProfile,
	config.NewTest,
	istiooperator.NewClient,
	k8s.NewAppsV1Client,
	k8s.NewClientset,
	k8s.NewConfig,
	k8s.NewDynamicClient,
	wire.Bind(new(dynamic.Interface), new(*dynamic.DynamicClient)),
	wire.Struct(new(http.Client)),
	wire.Struct(new(Test), "*"),
)

type Test struct {
	Config        *config.Test
	K8s           typedappsv1.AppsV1Interface
	IstioOperator *istiooperator.Client
}
