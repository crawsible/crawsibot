package irc_test

import (
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"
	"github.com/crawsible/crawsibot/irc/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forwarder", func() {
	var (
		fakeReceiver *mocks.FakeMsgRcvr
		forwarder    *irc.MessageForwarder
	)

	BeforeEach(func() {
		fakeReceiver = &mocks.FakeMsgRcvr{}
		forwarder = &irc.MessageForwarder{make(map[string][]irc.MsgRcvr)}
	})

	Describe("#EnrollForMsgs", func() {
		It("adds the receiver to its receiver list for the given command", func() {
			forwarder.EnrollForMsgs("SOMECMD")
			Expect(forwarder.MsgRcvrs["SOMECMD"]).To(ContainElement(fakeReceiver))
		})

		It("is idempotent", func() {
			forwarder.EnrollForMsgs("SOMECMD")
			receivers := forwarder.MsgRcvrs["SOMECMD"]

			forwarder.EnrollForMsgs("SOMECMD")
			Expect(forwarder.MsgRcvrs["SOMECMD"]).To(Equal(receivers))
		})
	})

	Describe("#StartForwarding", func() {
		var (
			fakeReader *mocks.FakeReader
			rdStrCh    chan string
			fakeCipher *mocks.FakeCipher
		)

		BeforeEach(func() {
			rdStrCh = make(chan string, 1)
			fakeReader = &mocks.FakeReader{RdStrCh: rdStrCh}
			fakeCipher = &mocks.FakeCipher{}
		})

		AfterEach(func() {
			rdStrCh <- "EOF"
		})

		It("reads from the provided ReadStringer, splitting on \\n", func() {
			rdStrCh <- ""
			forwarder.StartForwarding(fakeReader, fakeCipher)

			Eventually(fakeReader.ReadStringCalls).Should(Equal(1))
			Expect(fakeReader.ReadStringByte).Should(Equal(byte('\n')))
		})

		It("uses its cipher to decode the read messages", func() {
			rdStrCh <- "SOMECMD :some.server\r\n"
			forwarder.StartForwarding(fakeReader, fakeCipher)

			Eventually(fakeCipher.DecodeCalls).Should(Equal(1))
			Expect(fakeCipher.DecodeStrings[0]).To(Equal("SOMECMD :some.server\r\n"))
		})

		Context("with recipients", func() {
			var msg *models.Message
			BeforeEach(func() {
				forwarder.MsgRcvrs["SOMECMD"] = []irc.MsgRcvr{fakeReceiver}
				msg = &models.Message{
					Command: "SOMECMD",
					Params:  "some.server",
				}
				fakeCipher.DecodeMessages = []*models.Message{msg, msg}
			})

			It("calls RcvMsg with decoded message's field values on recipients", func() {
				rdStrCh <- ""
				forwarder.StartForwarding(fakeReader, fakeCipher)

				Eventually(fakeReceiver.RcvMsgCalls).Should(Equal(1))
				Expect(fakeReceiver.RcvMsgNick).To(BeZero())
				Expect(fakeReceiver.RcvMsgFprms).To(BeZero())
				Expect(fakeReceiver.RcvMsgPrms).To(Equal("some.server"))
			})

			It("works with multiple recipients", func() {
				secondRecipient := &mocks.FakeMsgRcvr{}
				forwarder.MsgRcvrs["SOMECMD"] = append(forwarder.MsgRcvrs["SOMECMD"], secondRecipient)
				rdStrCh <- ""
				forwarder.StartForwarding(fakeReader, fakeCipher)

				Eventually(fakeReceiver.RcvMsgCalls).Should(Equal(1))
				Expect(fakeReceiver.RcvMsgNick).To(BeZero())
				Expect(fakeReceiver.RcvMsgFprms).To(BeZero())
				Expect(fakeReceiver.RcvMsgPrms).To(Equal("some.server"))

				Eventually(secondRecipient.RcvMsgCalls).Should(Equal(1))
				Expect(secondRecipient.RcvMsgNick).To(BeZero())
				Expect(secondRecipient.RcvMsgFprms).To(BeZero())
				Expect(secondRecipient.RcvMsgPrms).To(Equal("some.server"))
			})

			It("handles multiple messages concurrently", func() {
				forwarder.StartForwarding(fakeReader, fakeCipher)
				rdStrCh <- ""
				rdStrCh <- ""

				Eventually(fakeReceiver.RcvMsgCalls).Should(Equal(2))
			})
		})
	})
})
