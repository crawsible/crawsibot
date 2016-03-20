package irc_test

import (
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"
	"github.com/crawsible/crawsibot/irc/message"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sender", func() {
	var (
		fakeWriter *mocks.FakeWriter
		fakeCipher *mocks.FakeCipher

		sender *irc.MessageSender
	)

	BeforeEach(func() {
		fakeWriter = &mocks.FakeWriter{}
		fakeCipher = &mocks.FakeCipher{}

		sender = &irc.MessageSender{}
	})

	Describe("#StartSending", func() {
		BeforeEach(func() {
			sender.StartSending(fakeWriter, fakeCipher)
			Eventually(sender.SendCh).ShouldNot(BeNil())
		})

		AfterEach(func() {
			close(sender.SendCh)
		})

		It("sets a channel with a buffer capacity of 90 as a field", func() {
			Expect(sender.SendCh).To(BeAssignableToTypeOf(make(chan *message.Message)))
			Expect(cap(sender.SendCh)).To(Equal(90))
		})

		It("uses its cipher to encode messages received on its chan", func() {
			sentMsg := &message.Message{
				Command:     "SOMECMD",
				FirstParams: "someparam",
			}
			sender.SendCh <- sentMsg

			Eventually(fakeCipher.EncodeCalls).Should(Equal(1))
			Expect(fakeCipher.EncodeMessages[0]).To(Equal(sentMsg))
		})

		It("writes to the provided conn with the encoded message", func() {
			fakeCipher.EncodeStrings = []string{"SOME encodedstring\r\n"}

			sender.SendCh <- &message.Message{}
			Eventually(fakeWriter.WriteCalls).Should(Equal(1))
			Expect(fakeWriter.WriteMessage).To(Equal([]byte("SOME encodedstring\r\n")))
		})
	})

	Describe("#Send", func() {
		var fakeCh chan *message.Message

		BeforeEach(func() {
			fakeCh = make(chan *message.Message, 1)
			sender.SendCh = fakeCh
		})

		AfterEach(func() {
			close(fakeCh)
		})

		It("converts args to messages and sends them through its chan", func() {
			sender.Send("SOMECMD", "some-fprms", "some-prms")
			Eventually(fakeCh).Should(HaveLen(1))
			expectedMsg := &message.Message{"", "SOMECMD", "some-fprms", "some-prms"}
			Expect(<-fakeCh).To(Equal(expectedMsg))
		})
	})
})
