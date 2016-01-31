package irc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIRC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IRC Suite")
}
