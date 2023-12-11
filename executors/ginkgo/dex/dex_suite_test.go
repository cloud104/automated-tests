package dex_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	typedappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"

	"github.com/cloud104/automated-tests/executors/ginkgo/dex/internal/config"
	"github.com/cloud104/automated-tests/executors/ginkgo/dex/internal/k8s"
)

var x = &struct {
	Config *config.Test
	K8s    typedappsv1.AppsV1Interface
}{}

var _ = BeforeSuite(func() {
	// Create a new Test configuration.
	test, err := config.NewTest()
	Expect(err).NotTo(HaveOccurred())

	// Create a new Profile configuration.
	profile, err := config.NewProfile()
	Expect(err).NotTo(HaveOccurred())

	// Create a Kubernetes REST configuration based on the provided profile.
	restConfig, err := k8s.NewConfig(profile)
	Expect(err).NotTo(HaveOccurred())

	// Create a Kubernetes clientset using the REST configuration.
	clientset, err := k8s.NewClientset(restConfig)
	Expect(err).NotTo(HaveOccurred())

	// Create an AppsV1 interface for interacting with Kubernetes applications.
	appsV1Interface := k8s.NewAppsV1Client(clientset)

	// Assign the created configurations and interfaces to the testing struct.
	x.Config = test
	x.K8s = appsV1Interface
})

func TestDex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dex Suite")
}
