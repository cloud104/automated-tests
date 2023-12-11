package dex_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/cloud104/automated-tests/executors/ginkgo/dex/internal/k8s"
)

var _ = Describe("Dex", Ordered, func() {
	When("the addon is installed", func() {
		It("should have a ready deployment for usage", func(ctx SpecContext) {
			Eventually(func() bool {
				// Define a label selector to identify the deployment
				labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
					"app.kubernetes.io/name":     "dex",
					"app.kubernetes.io/instance": "dex",
				}}

				// List deployments using the label selector
				list, err := x.K8s.Deployments(x.Config.Namespace).List(ctx, metav1.ListOptions{
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
			}).WithTimeout(x.Config.Timeout).Should(BeTrue())
		})
	})
})
