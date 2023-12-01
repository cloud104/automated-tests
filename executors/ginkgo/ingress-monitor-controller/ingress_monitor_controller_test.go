package ingress_monitor_controller_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/cloud104/automated-tests/executors/ginkgo/ingress-monitor-controller/internal/k8s"
	"github.com/cloud104/automated-tests/executors/ginkgo/ingress-monitor-controller/internal/wire"
)

var _ = Describe("IngressMonitorController", Ordered, func() {
	var test *wire.Test

	BeforeAll(func(ctx SpecContext) {
		var err error
		test, err = wire.SetUp(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	When("the addon is installed", func() {
		It("should have a ready deployment for usage", func(ctx SpecContext) {
			Eventually(func() bool {
				// Define a label selector to identify the deployment
				labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
					"app.kubernetes.io/managed-by": "Helm",
					"app.kubernetes.io/name":       "ingressmonitorcontroller",
				}}

				// List deployments using the label selector
				list, err := test.K8s.Deployments(test.Config.Namespace).List(ctx, metav1.ListOptions{
					LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
				})
				Expect(err).NotTo(HaveOccurred())

				// If no deployments are found, return false
				if len(list.Items) == 0 {
					return false
				}

				// Check if all deployments are available
				for _, deployment := range list.Items {
					if !k8s.IsDeploymentStatusConditionTrue(deployment.Status.Conditions, appsv1.DeploymentAvailable) {
						return false
					}
				}

				// If all deployments are available, return true
				return true
			}).WithTimeout(test.Config.Timeout).Should(BeTrue())
		})
	})
})
