package node_local_dns_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNodeLocalDns(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NodeLocalDns Suite")
}
