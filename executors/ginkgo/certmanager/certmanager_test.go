package certmanager_test

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/hashicorp/go-multierror"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/totvs-cloud/go-manifest"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	manifests  *manifest.Reader
	cluster_id = os.Getenv("CLUSTER_ID")
	region_dns = os.Getenv("REGION_DNS")
)

var _ = BeforeSuite(func() {
	config, err := loadRestConfig()
	Expect(err).NotTo(HaveOccurred())

	mr, err := manifest.NewReader("totvs-cloud", config)
	Expect(err).NotTo(HaveOccurred())

	manifests = mr
})

var _ = AfterSuite(func() {
	m, err := manifests.FromPath("./certificate.yaml", false)
	Expect(err).NotTo(HaveOccurred())

	err = m.Delete(context.Background())
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("Certmanager Test", func() {
	It("should return success apply", func() {

		sed("CLUSTERID", cluster_id, "./certificate.yaml")
		sed("REGIONDNS", region_dns, "./certificate.yaml")

		config, err := loadRestConfig()
		Expect(err).NotTo(HaveOccurred())

		// Instantiate a new ManifestReader by specifying the field manager and the Kubernetes cluster configuration
		mr, err := manifest.NewReader("totvs-cloud", config)
		Expect(err).NotTo(HaveOccurred())

		// Create a new Manifest object
		m, err := mr.FromPath("./certificate.yaml", false)
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
	if err1 == nil {
		return config, nil
	}

	result = multierror.Append(result, err1)

	// If that doesn't work, try getting it locally using the kubeconfig file in your home directory
	homeDir, err2 := os.UserHomeDir()
	if err2 != nil {
		return nil, multierror.Append(result, err2)
	}

	config, err3 := clientcmd.BuildConfigFromFlags("", path.Join(homeDir, "/.kube/config"))
	if err3 != nil {
		return nil, multierror.Append(result, err3)
	}

	return config, nil
}

func sed(old, new, filePath string) error {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	fileString := string(fileData)
	fileString = strings.ReplaceAll(fileString, old, new)
	fileData = []byte(fileString)

	err = ioutil.WriteFile(filePath, fileData, 0o600)
	if err != nil {
		return err
	}

	return nil
}
