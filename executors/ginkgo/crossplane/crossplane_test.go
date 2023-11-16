package crossplane_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/totvs-cloud/go-manifest"

	"github.com/cloud104/automated-tests/executors/ginkgo/crossplane/internal/wire"
)

const namespaceBasename = "crossplane-test"

var _ = Describe("Crossplane", func() {
	var (
		test    *wire.Test
		cleanup func()
	)
	var manifests manifest.List

	BeforeAll(func(ctx SpecContext) {
		var err error
		test, cleanup, err = wire.SetUp(ctx, namespaceBasename)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterAll(func(ctx SpecContext) {
		if test.Config.SkipDelete {
			return
		}

		defer cleanup()
		err := manifests.Delete(ctx)
		Expect(err).NotTo(HaveOccurred())
	})

	When("", func() {
		It("", func() {
			// TODO: implement tests scenarios
		})
	})
})
