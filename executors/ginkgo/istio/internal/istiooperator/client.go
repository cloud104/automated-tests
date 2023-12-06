package istiooperator

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	"github.com/cloud104/automated-tests/executors/ginkgo/istio/internal/config"
)

type Client struct {
	dynamicClient dynamic.Interface
	config        *config.Test
}

func NewClient(dynamicClient dynamic.Interface, config *config.Test) *Client {
	return &Client{dynamicClient: dynamicClient, config: config}
}

func (c *Client) GetResourceStatus(ctx context.Context) (*ResourceStatus, error) {
	var (
		resource  = schema.GroupVersionResource{Group: "install.istio.io", Version: "v1alpha1", Resource: "istiooperators"}
		name      = "tks-istio"
		namespace = c.config.Namespace
	)

	u, err := c.dynamicClient.Resource(resource).Namespace(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve IstioOperator resource '%s' in namespace '%s': %w", name, namespace, err)
	}

	data, err := u.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON for IstioOperator resource: %w", err)
	}

	rsrc, err := UnmarshalResource(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON for IstioOperator resource: %w", err)
	}

	return &rsrc.Status, nil
}
