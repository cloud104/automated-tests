package crossplane_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/totvs-cloud/go-manifest"
	"k8s.io/apimachinery/pkg/api/meta"

	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/wire"
)

const namespaceBasename = "crossplane-test"

var _ = Describe("Crossplane", Ordered, func() {
	var (
		test    *wire.Test
		cleanup func()
	)
	manifests := manifest.EmptyList()

	BeforeAll(func(ctx SpecContext) {
		var err error
		test, cleanup, err = wire.SetUp(ctx, namespaceBasename)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterAll(func(ctx SpecContext) {
		if test.Config.SkipDelete {
			return
		}

		err := test.Crossplane.DeleteManifests(ctx, manifests)
		Expect(err).NotTo(HaveOccurred())

		cleanup()
	})

	When("installing Kubernetes Provider", func() {
		var currentRevision string

		It("should apply Provider manifests", func(ctx SpecContext) {
			m, err := test.Crossplane.ApplyProviderManifests(ctx)
			Expect(err).NotTo(HaveOccurred())

			manifests = manifests.Append(m)
		})

		It("should eventually be healthy", func(ctx SpecContext) {
			Eventually(func() bool {
				status, err := test.Crossplane.GetProviderStatus(ctx)
				Expect(err).NotTo(HaveOccurred())

				if status == nil {
					return false
				}

				currentRevision = status.CurrentRevision
				return meta.IsStatusConditionTrue(status.Conditions, "Healthy")
			}).WithTimeout(test.Config.Timeout).Should(BeTrue())
		})

		It("should apply RBAC manifests", func(ctx SpecContext) {
			m, err := test.Crossplane.ApplyRBACManifests(ctx, currentRevision)
			Expect(err).NotTo(HaveOccurred())

			manifests = manifests.Append(m)
		})

		It("should apply ProviderConfig manifests", func(ctx SpecContext) {
			m, err := test.Crossplane.ApplyProviderConfigManifests(ctx)
			Expect(err).NotTo(HaveOccurred())

			manifests = manifests.Append(m)
		})
	})

	When("Kubernetes Provider is installed", func() {
		It("should create Kubernetes resources via Crossplane", func(ctx SpecContext) {
			m, err := test.Crossplane.ApplyObjectManifests(ctx)
			Expect(err).NotTo(HaveOccurred())

			manifests = manifests.Append(m)
		})

		It("should eventually have 1 ready pod", func(ctx SpecContext) {
			Eventually(func() int {
				count, err := test.Crossplane.CountReadySamplePods(ctx)
				Expect(err).NotTo(HaveOccurred())

				return count
			}).WithTimeout(test.Config.Timeout).Should(Equal(1))
		})
	})
})
