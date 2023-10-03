package externaldns_test

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/hashicorp/go-multierror"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/totvs-cloud/go-manifest"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var _ = Describe("deployment_manifest", func() {

	Context("Create resources through yaml manifest.", func() {
		config, err := loadRestConfig()
		Expect(err).NotTo(HaveOccurred())

		// Instantiate a new ManifestReader by specifying the field manager and the Kubernetes cluster configuration
		mr, err := manifest.NewReader("totvs-cloud", config)
		if err != nil {
			log.Fatal(err)
		}

		// Create a new Manifest object
		m, err := mr.FromPath("./manifest.yaml", false)
		if err != nil {
			log.Fatal(err)
		}

		// Apply the manifest using Server-Side Apply
		if err = m.Apply(context.Background()); err != nil {
			log.Fatal(err)
		}

		log.Println("Manifest applied successfully!")

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

	config, err3 := clientcmd.BuildConfigFromFlags("", path.Join(homeDir, "~/.kube/config"))
	if err3 != nil {
		return nil, multierror.Append(result, err3)
	}

	return config, nil
}
