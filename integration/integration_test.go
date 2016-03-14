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
			"-a", "localhost:3000",
			"-n", "someusername",
			"-p", "somepassword",
			"-c", "somechannel",
		)

		var err error
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		session.Terminate().Wait()
		for i := 0; i < len(reqCh); i++ {
			<-reqCh
		}
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("validates with the server using the specified credentials", func() {
		Eventually(reqCh).Should(Receive(Equal("PASS somepassword\r\n")))
		Eventually(reqCh).Should(Receive(Equal("NICK someusername\r\n")))
	})

	It("registers IRCv3 capabilities with the server", func() {
		Eventually(reqCh).Should(Receive(Equal("CAP REQ :twitch.tv/membership\r\n")))
	})

	It("PONGs when it gets PINGed", func() {
		Eventually(reqCh).Should(Receive(Equal("CAP REQ :twitch.tv/membership\r\n")))
		resCh <- "PING :tmi.twitch.tv\r\n"
		Eventually(reqCh).Should(Receive(Equal("PONG :tmi.twitch.tv\r\n")))
	})

	It("joins the specified channel", func() {
		Eventually(reqCh).Should(Receive(Equal("CAP REQ :twitch.tv/membership\r\n")))
		Consistently(reqCh).ShouldNot(Receive(Equal("JOIN #somechannel\r\n")))

		resCh <- ":tmi.twitch.tv 376 crawsibot :>\r\n"
		Eventually(reqCh).Should(Receive(Equal("JOIN #somechannel\r\n")))
	})

	It("Announces its arrival", func() {
		Eventually(reqCh).Should(Receive(Equal("CAP REQ :twitch.tv/membership\r\n")))
		resCh <- ":tmi.twitch.tv 376 crawsibot :>\r\n"
		Eventually(reqCh).Should(Receive(Equal("JOIN #somechannel\r\n")))

		resCh <- ":someusername.tmi.twitch.tv 366 someusername #somechannel :End of /NAMES list\r\n"
		Eventually(reqCh).Should(Receive(Equal("PRIVMSG #somechannel :COME WITH ME IF YOU WANT TO LIVE.")))
	})
})
