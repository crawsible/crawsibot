package irc_test

import (
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sender", func() {
	Describe("#StartSending", func() {
		var (
			fakeWriter *mocks.FakeWriter
			fakeCipher *mocks.FakeCipher

			sender *irc.Sender
			sendCh chan *irc.Message
		)

		BeforeEach(func() {
			fakeWriter = &mocks.FakeWriter{}
			fakeCipher = &mocks.FakeCipher{}

			sender = &irc.Sender{
				Encoder: fakeCipher,
			}
			sendCh = sender.StartSending(fakeWriter)
			Eventually(sendCh).ShouldNot(BeNil())
		})

		AfterEach(func() {
			close(sendCh)
		})

		It("returns a channel with a buffer capacity of 90", func() {
			Expect(cap(sendCh)).To(Equal(90))
		})

		It("uses its cipher to encode messages received on its chan", func() {
			sentMsg := &irc.Message{
				Command:     "SOMECMD",
				FirstParams: "someparam",
			}
			sendCh <- sentMsg

			Eventually(fakeCipher.EncodeCalls).Should(Equal(1))
			Expect(fakeCipher.EncodeMessages[0]).To(Equal(sentMsg))
		})

		It("writes to the provided conn with the encoded message", func() {
			fakeCipher.EncodeReturns = []string{"SOME encodedstring\r\n"}

			sendCh <- &irc.Message{}
			Eventually(fakeWriter.WriteCalls).Should(Equal(1))
			Expect(fakeWriter.WriteMessage).To(Equal([]byte("SOME encodedstring\r\n")))
		})
	})
})
