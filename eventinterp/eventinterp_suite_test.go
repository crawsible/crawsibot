package eventinterp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEventInterp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EventInterp Suite")
}
