package irc_test

import (
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"
	"github.com/crawsible/crawsibot/irc/message"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ponger", func() {
	var ponger *irc.ServerPonger

	Describe("#StartPonging", func() {
		var (
			pingCh   chan *message.Message
			fakeMsgr *mocks.FakeMessenger
		)

		BeforeEach(func() {
			pingCh = make(chan *message.Message, 1)
			fakeMsgr = &mocks.FakeMessenger{EnrollForMsgsReturnChan: pingCh}
			ponger = &irc.ServerPonger{}

			ponger.StartPonging(fakeMsgr)
		})

		It("sets its PingCh with the chan provided by its Messenger", func() {
			Expect(fakeMsgr.EnrollForMsgsCalls).To(Equal(1))
			Expect(fakeMsgr.EnrollForMsgsCmd).To(Equal("PING"))

			Expect(ponger.PingCh).NotTo(Receive())
			pingCh <- &message.Message{}
			Expect(ponger.PingCh).To(Receive())
		})

		Context("once ponging has started", func() {
			It("PONGs the received server with the messenger", func() {
				ponger.PingCh <- &message.Message{Params: "some.server"}

				Eventually(fakeMsgr.SendCalls).Should(Equal(1))
				expectedArgs := []string{"PONG", "", "some.server"}
				Expect(fakeMsgr.SendArgs[0]).To(Equal(expectedArgs))
			})

			It("handles multiple messages", func() {
				ponger.PingCh <- &message.Message{Params: "some.server"}
				ponger.PingCh <- &message.Message{Params: "some.other.server"}

				Eventually(fakeMsgr.SendCalls).Should(Equal(2))
				expectedArgs0 := []string{"PONG", "", "some.server"}
				expectedArgs1 := []string{"PONG", "", "some.other.server"}
				Expect(fakeMsgr.SendArgs[0]).To(Equal(expectedArgs0))
				Expect(fakeMsgr.SendArgs[1]).To(Equal(expectedArgs1))
			})
		})
	})
})
