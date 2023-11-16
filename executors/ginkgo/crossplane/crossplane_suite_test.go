package crossplane_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCrossplane(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crossplane Suite")
}
