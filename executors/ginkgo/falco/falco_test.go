package falco_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/cloud104/automated-tests/executors/ginkgo/falco/internal/k8s"
	"github.com/cloud104/automated-tests/executors/ginkgo/falco/internal/wire"
)

var _ = Describe("Falco", Ordered, func() {
	var test *wire.Test

	BeforeAll(func(ctx SpecContext) {
		var err error
		test, err = wire.SetUp(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	When("the addon is installed", func() {
		Context("with the falco", func() {
			It("should have a DaemonSet", func(ctx SpecContext) {
				Eventually(func() int {
					// Fetches the list of Falco DaemonSet
					labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
						"app.kubernetes.io/instance": "falco",
						"app.kubernetes.io/name":     "falco",
					}}
					daemonsets, err := test.AppsV1.DaemonSets(test.Config.Namespace).List(ctx, metav1.ListOptions{
						LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
					})
					Expect(err).NotTo(HaveOccurred())

					// Return the number of DaemonSets found
					return len(daemonsets.Items)
				}).WithTimeout(test.Config.Timeout).Should(Equal(1))
			})

			It("should have a healthy pod in each node", func(ctx SpecContext) {
				Eventually(func() bool {
					// Fetches the list of Kubernetes nodes
					nodes, err := test.CoreV1.Nodes().List(ctx, metav1.ListOptions{})
					Expect(err).NotTo(HaveOccurred())

					// Initializes a map to store the health status of each Kubernetes node
					nodeHealth := make(map[string]bool)
					for _, node := range nodes.Items {
						nodeHealth[node.GetName()] = false
					}

					// Fetches the list of Falco pods
					labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
						"app.kubernetes.io/instance": "falco",
						"app.kubernetes.io/name":     "falco",
					}}
					pods, err := test.CoreV1.Pods(test.Config.Namespace).List(ctx, metav1.ListOptions{
						LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
					})
					Expect(err).NotTo(HaveOccurred())

					// Updates the node health status based on the readiness of the pods running on each node
					for _, pod := range pods.Items {
						nodeHealth[pod.Spec.NodeName] = k8s.IsPodStatusConditionTrue(pod.Status.Conditions, corev1.ContainersReady)
					}

					// Checks if there is any node without a health status
					for _, health := range nodeHealth {
						if !health {
							return false
						}
					}

					return true
				}).WithTimeout(test.Config.Timeout).Should(BeTrue())
			})
		})

		Context("with the falco-exporter", func() {
			It("should have a DaemonSet", func(ctx SpecContext) {
				Eventually(func() int {
					// Fetches the list of Falco DaemonSet
					labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
						"app.kubernetes.io/instance": "falco-exporter",
						"app.kubernetes.io/name":     "falco-exporter",
					}}
					daemonsets, err := test.AppsV1.DaemonSets(test.Config.Namespace).List(ctx, metav1.ListOptions{
						LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
					})
					Expect(err).NotTo(HaveOccurred())

					// Return the number of DaemonSets found
					return len(daemonsets.Items)
				}).WithTimeout(test.Config.Timeout).Should(Equal(1))
			})

			It("should have a healthy pod in each node", func(ctx SpecContext) {
				Eventually(func() bool {
					// Fetches the list of Kubernetes nodes
					nodes, err := test.CoreV1.Nodes().List(ctx, metav1.ListOptions{})
					Expect(err).NotTo(HaveOccurred())

					// Initializes a map to store the health status of each Kubernetes node
					nodeHealth := make(map[string]bool)
					for _, node := range nodes.Items {
						nodeHealth[node.GetName()] = false
					}

					// Fetches the list of FalcoExporter pods
					labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{
						"app.kubernetes.io/instance": "falco-exporter",
						"app.kubernetes.io/name":     "falco-exporter",
					}}
					pods, err := test.CoreV1.Pods(test.Config.Namespace).List(ctx, metav1.ListOptions{
						LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
					})
					Expect(err).NotTo(HaveOccurred())

					// Updates the node health status based on the readiness of the pods running on each node
					for _, pod := range pods.Items {
						nodeHealth[pod.Spec.NodeName] = k8s.IsPodStatusConditionTrue(pod.Status.Conditions, corev1.ContainersReady)
					}

					// Checks if there is any node without a health status
					for _, health := range nodeHealth {
						if !health {
							return false
						}
					}

					return true
				}).WithTimeout(test.Config.Timeout).Should(BeTrue())
			})
		})
	})
})
