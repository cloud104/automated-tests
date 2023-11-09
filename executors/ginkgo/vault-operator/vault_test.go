package vault_operator_test

import (
	_ "embed"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/totvs-cloud/go-manifest"

	"github.com/cloud104/automated-tests/executors/ginkgo/vault-operator/internal/wire"
)

var _ = Describe("Installing Vault using the Operator", Ordered, func() {
	var (
		test      *wire.Test
		namespace string
		manifests manifest.List
	)

	BeforeAll(func(ctx SpecContext) {
		t, err := wire.SetUp(ctx)
		Expect(err).NotTo(HaveOccurred())

		ns, err := t.TemporaryNamespace.Create(ctx, "vault-operator-test")
		Expect(err).NotTo(HaveOccurred())

		test = t
		namespace = ns
	})

	AfterAll(func(ctx SpecContext) {
		err := test.VaultOperator.CleanUp(ctx, namespace, manifests)
		Expect(err).NotTo(HaveOccurred())

		err = test.TemporaryNamespace.CleanUp(ctx, namespace)
		Expect(err).NotTo(HaveOccurred())
	})

	When("the Operator CRD is applied", func() {
		BeforeAll(func(ctx SpecContext) {
			m, err := test.VaultOperator.ApplyManifests(ctx, namespace)
			Expect(err).NotTo(HaveOccurred())

			manifests = m
		})

		It("should eventually have 1 ready pod", func(ctx SpecContext) {
			countReadyPods := func() int {
				count, err := test.VaultOperator.CountReadyPods(ctx, namespace)
				Expect(err).NotTo(HaveOccurred())

				return count
			}
			Eventually(countReadyPods, ctx).Should(Equal(1))
		}, SpecTimeout(time.Second*30))
	})

	When("the installation process is complete", func() {
		It("should be able to add a new secret to Vault", func(ctx SpecContext) {
			req := schema.KvV2WriteRequest{Data: map[string]any{"password": "baboseiras"}}
			_, err := test.VaultSecrets.KvV2Write(ctx, "robopato", req, vault.WithMountPath("secret"))
			Expect(err).NotTo(HaveOccurred())
		})

		It("should be able to read the new secret from Vault", func(ctx SpecContext) {
			s, err := test.VaultSecrets.KvV2Read(ctx, "robopato", vault.WithMountPath("secret"))
			Expect(err).NotTo(HaveOccurred())

			password := s.Data.Data["password"]
			Expect(password).To(Equal("baboseiras"))
		})
	})
})
