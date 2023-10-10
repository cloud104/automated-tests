package externaldns_ingress_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

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
	m, err := manifests.FromPath("./resources.yaml", false)
	Expect(err).NotTo(HaveOccurred())

	err = m.Delete(context.Background())
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("Application Test", func() {

	BeforeEach(func() {

		sed("CLUSTERID", cluster_id, "./resources.yaml")
		sed("REGIONDNS", region_dns, "./resources.yaml")

		config, err := loadRestConfig()
		Expect(err).NotTo(HaveOccurred())

		mr, err := manifest.NewReader("totvs-cloud", config)
		Expect(err).NotTo(HaveOccurred())

		// Create a new Manifest object
		m, err := mr.FromPath("./resources.yaml", false)
		Expect(err).NotTo(HaveOccurred())

		// Apply the manifest using Server-Side Apply
		err = m.Apply(context.Background())
		Expect(err).NotTo(HaveOccurred())

		time.Sleep(10 * time.Second)

		GinkgoWriter.Println("Manifest applied successfully!")
	})

	It("should return status code 200 when request Ingress", func() {

		baseUrl, err := url.Parse("https://check-ingress." + cluster_id + "." + region_dns)
		Expect(err).NotTo(HaveOccurred())

		resp, requestErr := http.Get(baseUrl.String())

		Expect(requestErr).To(BeNil())
		Expect(resp.StatusCode).To(Equal(200))

		GinkgoWriter.Println(baseUrl)
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
