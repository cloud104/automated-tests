package falco_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFalco(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Falco Suite")
}
