package integration_test

import (
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
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

	FIt("complains if it's missing flags", func() {
		command := exec.Command(
			crawsibotPath,
			"-a localhost:3000",
			"-n some-username",
		)

		var err error
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session.Err).Should(gbytes.Say(fmt.Sprintf("Usage of %s:", crawsibotPath)))
		Expect(session.Err.Contents()).NotTo(ContainSubstring("flag provided but not defined"))
	})

	Describe("with all necessary arguments", func() {
		BeforeEach(func() {
			command := exec.Command(
				crawsibotPath,
				"-a localhost:3000",
				"-n some-username",
				"-p some-password",
				"-c somechannel",
			)

			var err error
			session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("validates with the server using the specified credentials", func() {
			Eventually(reqCh).Should(Receive(Equal("PASS some-password\r\n")))
			Eventually(reqCh).Should(Receive(Equal("NICK some-username\r\n")))
		})

		XIt("joins the specified channel", func() {
			Eventually(reqCh).Should(Receive(Equal("JOIN #somechannel\r\n")))
		})

		XIt("Announces its arrival", func() {
			Eventually(reqCh).Should(Receive(Equal("PRIVMSG #somechannel :COME WITH ME IF YOU WANT TO LIVE.")))
		})

		XIt("PONGs when it gets PINGED", func() {})
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})
})
