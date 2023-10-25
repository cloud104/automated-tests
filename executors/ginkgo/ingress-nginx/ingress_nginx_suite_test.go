package ingress_nginx_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIngressNginx(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IngressNginx Suite")
}
