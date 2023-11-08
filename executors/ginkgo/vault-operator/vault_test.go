package vault_operator_test

import (
	_ "embed"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/totvs-cloud/go-manifest"
)

var _ = Describe("Installing Vault using the Operator", Ordered, func() {
	var (
		testing   *Testing
		namespace string
		manifests manifest.List
	)

	BeforeAll(func(ctx SpecContext) {
		t, err := initialize(ctx)
		Expect(err).NotTo(HaveOccurred())

		ns, err := t.TemporaryNamespace.Create(ctx, "vault-operator-test")
		Expect(err).NotTo(HaveOccurred())

		testing = t
		namespace = ns
	})

	AfterAll(func(ctx SpecContext) {
		err := testing.VaultOperator.CleanUp(ctx, namespace, manifests)
		Expect(err).NotTo(HaveOccurred())

		err = testing.TemporaryNamespace.CleanUp(ctx, namespace)
		Expect(err).NotTo(HaveOccurred())
	})

	When("the Operator CRD is applied", func() {
		BeforeAll(func(ctx SpecContext) {
			m, err := testing.VaultOperator.ApplyManifests(ctx, namespace)
			Expect(err).NotTo(HaveOccurred())

			manifests = m
		})

		It("should eventually have 1 ready pod", func(ctx SpecContext) {
			countReadyPods := func() int {
				count, err := testing.VaultOperator.CountReadyPods(ctx, namespace)
				Expect(err).NotTo(HaveOccurred())

				return count
			}
			Eventually(countReadyPods, ctx).Should(Equal(1))
		}, SpecTimeout(time.Second*30))
	})

	When("the installation process is complete", func() {
		It("should be able to add a new secret to Vault", func(ctx SpecContext) {
			req := schema.KvV2WriteRequest{Data: map[string]any{"password": "baboseiras"}}
			_, err := testing.VaultSecrets.KvV2Write(ctx, "robopato", req, vault.WithMountPath("secret"))
			Expect(err).NotTo(HaveOccurred())
		})

		It("should be able to read the new secret from Vault", func(ctx SpecContext) {
			s, err := testing.VaultSecrets.KvV2Read(ctx, "robopato", vault.WithMountPath("secret"))
			Expect(err).NotTo(HaveOccurred())

			password := s.Data.Data["password"]
			Expect(password).To(Equal("baboseiras"))
		})
	})
})
