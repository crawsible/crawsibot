package irc_test

import (
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forwarder", func() {
	var (
		fakeRecipient *mocks.FakeRecipient

		forwarder *irc.Forwarder
	)

	BeforeEach(func() {
		fakeRecipient = &mocks.FakeRecipient{}
		forwarder = &irc.Forwarder{}
	})

	Describe("#EnrollForPING", func() {
		BeforeEach(func() {
			forwarder.EnrollForPING(fakeRecipient)
		})

		It("adds argument to its list of PINGRecipients", func() {
			Expect(forwarder.PINGRecipients).To(ContainElement(fakeRecipient))
		})

		It("is idempotent", func() {
			recipients := forwarder.PINGRecipients
			forwarder.EnrollForPING(fakeRecipient)
			Expect(forwarder.PINGRecipients).To(Equal(recipients))
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
				forwarder.PINGRecipients = []irc.PINGRecipient{fakeRecipient}
				msg = &irc.Message{
					Command: "PING",
					Params:  "some.server",
				}
				fakeCipher.DecodeMessages = []*irc.Message{msg}
			})

			It("calls RcvPING with decoded message's field values on recipients", func() {
				rdStrCh <- ""
				forwarder.StartForwarding(fakeReader, fakeCipher)

				Eventually(fakeRecipient.RcvPINGCalls).Should(Equal(1))
				Expect(fakeRecipient.RcvPINGNick).To(BeZero())
				Expect(fakeRecipient.RcvPINGFprms).To(BeZero())
				Expect(fakeRecipient.RcvPINGPrms).To(Equal("some.server"))
			})

			It("works with multiple recipients", func() {
				secondRecipient := &mocks.FakeRecipient{}
				forwarder.PINGRecipients = append(forwarder.PINGRecipients, secondRecipient)
				rdStrCh <- ""
				forwarder.StartForwarding(fakeReader, fakeCipher)

				Eventually(fakeRecipient.RcvPINGCalls).Should(Equal(1))
				Expect(fakeRecipient.RcvPINGNick).To(BeZero())
				Expect(fakeRecipient.RcvPINGFprms).To(BeZero())
				Expect(fakeRecipient.RcvPINGPrms).To(Equal("some.server"))

				Eventually(secondRecipient.RcvPINGCalls).Should(Equal(1))
				Expect(secondRecipient.RcvPINGNick).To(BeZero())
				Expect(secondRecipient.RcvPINGFprms).To(BeZero())
				Expect(secondRecipient.RcvPINGPrms).To(Equal("some.server"))
			})

			It("handles multiple messages concurrently", func() {
				forwarder.StartForwarding(fakeReader, fakeCipher)
				rdStrCh <- ""
				rdStrCh <- ""

				Eventually(fakeRecipient.RcvPINGCalls).Should(Equal(2))
			})
		})
	})
})
