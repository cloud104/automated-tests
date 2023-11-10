package vault_operator_test

import (
	_ "embed"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/totvs-cloud/go-manifest"

	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/wire"
)

const namespaceBasename = "vault-operator-test"

var _ = Describe("Installing Vault using the Operator", Ordered, func() {
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

	When("the Operator CRD is applied", func() {
		It("should apply Vault manifests successfully", func(ctx SpecContext) {
			m, err := test.Operator.ApplyManifests(ctx)
			Expect(err).NotTo(HaveOccurred())

			manifests = m
		})

		It("should eventually have 1 ready pod", func(ctx SpecContext) {
			Eventually(func() int {
				count, err := test.Operator.CountReadyPods(ctx)
				Expect(err).NotTo(HaveOccurred())

				return count
			}).WithTimeout(test.Config.Timeout).Should(Equal(1))
		})
	})

	When("the installation process is complete", func() {
		It("should be able to add a new secret to Vault", func(ctx SpecContext) {
			Eventually(func() error {
				req := schema.KvV2WriteRequest{Data: map[string]any{"password": "baboseiras"}}
				_, err := test.Secrets.KvV2Write(ctx, "robopato", req, vault.WithMountPath("secret"))
				return err
			}).WithTimeout(test.Config.Timeout).Should(Succeed())
		})

		It("should be able to read the new secret from Vault", func(ctx SpecContext) {
			Eventually(func() any {
				s, err := test.Secrets.KvV2Read(ctx, "robopato", vault.WithMountPath("secret"))
				Expect(err).NotTo(HaveOccurred())

				return s.Data.Data["password"]
			}).WithTimeout(test.Config.Timeout).Should(Equal("baboseiras"))
		})
	})
})
