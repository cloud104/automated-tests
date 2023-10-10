package certmanager_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCertmanager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Certmanager Suite")
}
