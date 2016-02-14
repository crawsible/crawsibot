package irc_test

import (
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forwarder", func() {
	var (
		fakeReceiver *mocks.FakeReceiver

		forwarder *irc.MessageForwarder
	)

	BeforeEach(func() {
		fakeReceiver = &mocks.FakeReceiver{}
		forwarder = &irc.MessageForwarder{}
	})

	Describe("#EnrollForMsgs", func() {
		It("is idempotent", func() {
			forwarder.EnrollForMsgs(fakeReceiver, "PING")
			receivers := forwarder.PINGRcvrs

			forwarder.EnrollForMsgs(fakeReceiver, "PING")
			Expect(forwarder.PINGRcvrs).To(Equal(receivers))
		})

		Context("when called with 'PING'", func() {
			BeforeEach(func() {
				forwarder.EnrollForMsgs(fakeReceiver, "PING")
			})

			It("adds the argument to its list of PINGRecipients", func() {
				Expect(forwarder.PINGRcvrs).To(ContainElement(fakeReceiver))
			})
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
			rdStrCh <- "PING :some.server\r\n"
			forwarder.StartForwarding(fakeReader, fakeCipher)

			Eventually(fakeCipher.DecodeCalls).Should(Equal(1))
			Expect(fakeCipher.DecodeStrings[0]).To(Equal("PING :some.server\r\n"))
		})

		Context("with recipients", func() {
			var msg *irc.Message
			BeforeEach(func() {
				forwarder.PINGRcvrs = []irc.MsgRcvr{fakeReceiver}
				msg = &irc.Message{
					Command: "PING",
					Params:  "some.server",
				}
				fakeCipher.DecodeMessages = []*irc.Message{msg}
			})

			It("calls RcvPING with decoded message's field values on recipients", func() {
				rdStrCh <- ""
				forwarder.StartForwarding(fakeReader, fakeCipher)

				Eventually(fakeReceiver.RcvMsgCalls).Should(Equal(1))
				Expect(fakeReceiver.RcvMsgNick).To(BeZero())
				Expect(fakeReceiver.RcvMsgFprms).To(BeZero())
				Expect(fakeReceiver.RcvMsgPrms).To(Equal("some.server"))
			})

			It("works with multiple recipients", func() {
				secondRecipient := &mocks.FakeReceiver{}
				forwarder.PINGRcvrs = append(forwarder.PINGRcvrs, secondRecipient)
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
