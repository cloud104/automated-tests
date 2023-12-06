package istio_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/cloud104/automated-tests/executors/ginkgo/istio/internal/k8s"
	"github.com/cloud104/automated-tests/executors/ginkgo/istio/internal/wire"
)

var _ = Describe("Istio", Ordered, func() {
	var test *wire.Test

	BeforeAll(func(ctx SpecContext) {
		var err error
		test, err = wire.SetUp(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	When("the IstioOperator CRD is applied", func() {
		It("is expected that all components should be healthy", func(ctx SpecContext) {
			Eventually(func() bool {
				// Get the status of the IstioOperator CRD resource
				resource, err := test.IstioOperator.GetResourceStatus(ctx)
				Expect(err).NotTo(HaveOccurred())

				// Check if the Istio components are in a healthy state.
				return resource.Status == "HEALTHY" &&
					resource.ComponentStatus.Base.Status == "HEALTHY" &&
					resource.ComponentStatus.EgressGateways.Status == "HEALTHY" &&
					resource.ComponentStatus.IngressGateways.Status == "HEALTHY" &&
					resource.ComponentStatus.Pilot.Status == "HEALTHY"
			}).WithTimeout(test.Config.Timeout).Should(BeTrue())
		})
	})

	When("the addon is installed", func() {
		It("should have an available egressgateway deployment", func(ctx SpecContext) {
			Eventually(func() bool {
				// Define a label selector to identify the deployment
				labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
					"install.operator.istio.io/owning-resource":           "tks-istio",
					"install.operator.istio.io/owning-resource-namespace": "istio-system",
					"operator.istio.io/component":                         "EgressGateways",
					"app":                                                 "istio-egressgateway",
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

		It("should have an available ingressgateway deployment", func(ctx SpecContext) {
			Eventually(func() bool {
				// Define a label selector to identify the deployment
				labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
					"install.operator.istio.io/owning-resource":           "tks-istio",
					"install.operator.istio.io/owning-resource-namespace": "istio-system",
					"operator.istio.io/component":                         "IngressGateways",
					"app":                                                 "istio-ingressgateway",
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

		It("should have an available istiod deployment", func(ctx SpecContext) {
			Eventually(func() bool {
				// Define a label selector to identify the deployment
				labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
					"install.operator.istio.io/owning-resource":           "tks-istio",
					"install.operator.istio.io/owning-resource-namespace": "istio-system",
					"operator.istio.io/component":                         "Pilot",
					"app":                                                 "istiod",
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

		It("should have an available kiali deployment", func(ctx SpecContext) {
			Eventually(func() bool {
				// Define a label selector to identify the deployment
				labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
					"app":                          "kiali",
					"app.kubernetes.io/managed-by": "Helm",
					"app.kubernetes.io/name":       "kiali",
					"app.kubernetes.io/part-of":    "kiali",
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

		It("should have an available kiali-oauth2-proxy deployment", func(ctx SpecContext) {
			Eventually(func() bool {
				// Define a label selector to identify the deployment
				labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
					"k8s-app": "kiali-oauth2-proxy",
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

		It("should have an available prometheus deployment", func(ctx SpecContext) {
			Eventually(func() bool {
				// Define a label selector to identify the deployment
				labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
					"app": "prometheus",
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
