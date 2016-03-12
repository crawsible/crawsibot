package chatapp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestChatApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ChatApp Suite")
}
