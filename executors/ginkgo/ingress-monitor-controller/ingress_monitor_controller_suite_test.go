package ingress_monitor_controller_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIngressMonitorController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IngressMonitorController Suite")
}
