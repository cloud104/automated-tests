package externaldns_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExternaldns(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Externaldns Suite")
}
