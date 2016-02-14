package irc_test

import (
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ponger", func() {
	var ponger *irc.ServerPonger

	Describe("#StartPonging", func() {
		var fakeMsgr *mocks.FakeMessenger

		BeforeEach(func() {
			fakeMsgr = &mocks.FakeMessenger{}
			ponger = &irc.ServerPonger{}

			ponger.StartPonging(fakeMsgr)
		})

		It("sets its blocking ping listening channel", func() {
			ch := make(chan string)
			defer close(ch)

			Expect(ponger.PingCh).NotTo(BeNil())
			Expect(cap(ponger.PingCh)).To(Equal(0))
		})

		It("enrolls itself for PINGs with the messenger", func() {
			Expect(fakeMsgr.EnrollForPINGCalls).To(Equal(1))
			Expect(fakeMsgr.EnrollForPINGMsgRcvr).To(Equal(ponger))
		})

		Context("once ponging has started", func() {
			It("PONGs the received server with the messenger", func() {
				ponger.PingCh <- "some-server"

				Eventually(fakeMsgr.SendCalls).Should(Equal(1))
				expectedArgs := []string{"PONG", "", "some-server"}
				Expect(fakeMsgr.SendArgs[0]).To(Equal(expectedArgs))
			})

			It("handles multiple messages", func() {
				ponger.PingCh <- "some-server"
				ponger.PingCh <- "another-server"

				Eventually(fakeMsgr.SendCalls).Should(Equal(2))
				expectedArgs0 := []string{"PONG", "", "some-server"}
				expectedArgs1 := []string{"PONG", "", "another-server"}
				Expect(fakeMsgr.SendArgs[0]).To(Equal(expectedArgs0))
				Expect(fakeMsgr.SendArgs[1]).To(Equal(expectedArgs1))
			})
		})
	})

	Describe("#RcvMsg", func() {
		var fakeCh chan string

		BeforeEach(func() {
			fakeCh = make(chan string, 1)
			ponger = &irc.ServerPonger{PingCh: fakeCh}
		})

		AfterEach(func() {
			close(fakeCh)
		})

		It("sends a constructed PingMsg to its PingCh", func() {
			ponger.RcvMsg("", "", "some-server")

			Expect(fakeCh).To(HaveLen(1))
			Expect(<-fakeCh).To(Equal("some-server"))
		})
	})
})
