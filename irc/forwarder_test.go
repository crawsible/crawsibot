package irc_test

import (
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"
	"github.com/crawsible/crawsibot/irc/message"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Forwarder", func() {
	var forwarder *irc.MessageForwarder

	BeforeEach(func() {
		forwarder = &irc.MessageForwarder{make(map[string][]chan *message.Message)}
	})

	Describe("#EnrollForMsgs", func() {
		It("returns a new one-buffer Message chan", func() {
			newCh := forwarder.EnrollForMsgs("SOMECMD")

			var msgCh chan *message.Message
			Expect(newCh).To(BeAssignableToTypeOf(msgCh))
			Expect(cap(newCh)).To(Equal(1))
		})

		It("adds the chan to its slice of Message chans for the given cmd", func() {
			newCh := forwarder.EnrollForMsgs("SOMECMD")

			forwarder.MsgChs["SOMECMD"][0] <- &message.Message{}
			Expect(newCh).To(Receive())
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
			var (
				msg   *message.Message
				msgCh chan *message.Message
			)
			BeforeEach(func() {
				msgCh = make(chan *message.Message)
				forwarder.MsgChs["SOMECMD"] = []chan *message.Message{msgCh}
				msg = &message.Message{
					Command: "SOMECMD",
					Params:  "some.server",
				}
				fakeCipher.DecodeMessages = []*message.Message{msg, msg}
			})

			It("sends the decoded message on each of the appropriate channels", func() {
				rdStrCh <- ""
				forwarder.StartForwarding(fakeReader, fakeCipher)
				Eventually(msgCh).Should(Receive(Equal(msg)))
			})

			It("works with multiple recipients", func() {
				secondCh := make(chan *message.Message)
				forwarder.MsgChs["SOMECMD"] = append(forwarder.MsgChs["SOMECMD"], secondCh)
				rdStrCh <- ""
				forwarder.StartForwarding(fakeReader, fakeCipher)

				Eventually(msgCh).Should(Receive(Equal(msg)))
				Eventually(secondCh).Should(Receive(Equal(msg)))
			})
		})
	})
})
