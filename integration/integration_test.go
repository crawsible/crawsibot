package integration_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Integration", func() {
	var crawsibotPath string
	var session *gexec.Session

	BeforeSuite(func() {
		var err error
		crawsibotPath, err = gexec.Build("github.com/crawsible/crawsibot")
		Expect(err).NotTo(HaveOccurred())
	})

	BeforeEach(func() {
		command := exec.Command(
			crawsibotPath,
			"-a localhost:3000",
			"-u some-username",
			"-p some-password",
		)

		var err error
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	FIt("validates with the server", func() {
		Expect(<-reqCh).To(Equal("PASS some-password\r\n"))
		Expect(<-reqCh).To(Equal("NICK some-username\r\n"))
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})
})
