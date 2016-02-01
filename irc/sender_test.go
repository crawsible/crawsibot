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
			fakeConn   *mocks.FakeConn
			fakeCipher *mocks.FakeCipher

			sender *irc.Sender
			sendCh chan *irc.Message
		)

		BeforeEach(func() {
			fakeConn = &mocks.FakeConn{}
			fakeCipher = &mocks.FakeCipher{}

			sender = &irc.Sender{
				Encoder: fakeCipher,
			}
			sendCh = sender.StartSending(fakeConn)
			Eventually(sendCh).ShouldNot(BeNil())
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
			Eventually(fakeConn.WriteCalls).Should(Equal(1))
			Expect(fakeConn.WriteMessage).To(Equal([]byte("SOME encodedstring\r\n")))
		})

		AfterEach(func() {
			close(sendCh)
		})
	})
})
