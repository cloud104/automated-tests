package externaldns_ingress_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExternaldnsIngress(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ExternaldnsIngress Suite")
}
