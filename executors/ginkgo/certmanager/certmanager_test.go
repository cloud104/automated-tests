package certmanager_test

import (
	"context"

	"github.com/hashicorp/go-multierror"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/totvs-cloud/go-manifest"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var manifests *manifest.Reader

var _ = BeforeSuite(func() {
	config, err := loadRestConfig()
	Expect(err).NotTo(HaveOccurred())

	mr, err := manifest.NewReader("totvs-cloud", config)
	Expect(err).NotTo(HaveOccurred())

	manifests = mr
})

var _ = AfterSuite(func() {
	m, err := manifests.FromPath("./certificate-test.yaml", false)
	Expect(err).NotTo(HaveOccurred())

	err = m.Delete(context.Background())
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("Certmanager Test", func() {
	It("should return success apply", func() {
		config, err := loadRestConfig()
		Expect(err).NotTo(HaveOccurred())

		// Instantiate a new ManifestReader by specifying the field manager and the Kubernetes cluster configuration
		mr, err := manifest.NewReader("totvs-cloud", config)
		Expect(err).NotTo(HaveOccurred())

		// Create a new Manifest object
		m, err := mr.FromPath("./certificate-test.yaml", false)
		Expect(err).NotTo(HaveOccurred())

		// Apply the manifest using Server-Side Apply
		err = m.Apply(context.Background())
		Expect(err).NotTo(HaveOccurred())

		GinkgoWriter.Println("Manifest applied successfully!")
	})
})

func loadRestConfig() (*rest.Config, error) {
	var result error

	// Attempt to retrieve the Kubernetes cluster configuration first from the in-cluster environment
	config, err1 := rest.InClusterConfig()
	if err1 != nil {
		panic(err.Error())
	}

	replaceTransport(config, reqClient.GetTransport())

	// Create kubernetes client with config.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	result = multierror.Append(result, err1)

	// If that doesn't work, try getting it locally using the kubeconfig file in your home directory
	//	homeDir, err2 := os.UserHomeDir()
	//	if err2 != nil {
	//		return nil, multierror.Append(result, err2)
	//	}
	//
	//	config, err3 := clientcmd.BuildConfigFromFlags("", path.Join(homeDir, "/.kube/config"))
	//	if err3 != nil {
	//		return nil, multierror.Append(result, err3)
	//	}

	return clientset, nil
}

func replaceTransport(config *rest.Config, t *req.Transport) {
	// Extract tls.Config from rest.Config
	tlsConfig, err := rest.TLSConfigFor(config)
	if err != nil {
		panic(err.Error())
	}
	// Set TLSClientConfig to req's Transport.
	t.TLSClientConfig = tlsConfig
	// Override with req's Transport.
	config.Transport = t
	// rest.Config.TLSClientConfig should be empty if
	// custom Transport been set.
	config.TLSClientConfig = rest.TLSClientConfig{}
}
