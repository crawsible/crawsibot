package irc_test

import (
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/message"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MessageCipher", func() {
	var cipher *irc.MessageCipher

	Describe("#Decode", func() {
		It("parses IRC messages into component parts", func() {
			msg, err := cipher.Decode(":some-nick!some-user@some.host SOMECMD some :params\r\n")

			Expect(err).NotTo(HaveOccurred())
			Expect(msg.NickOrSrvname).To(Equal("some-nick"))
			Expect(msg.Command).To(Equal("SOMECMD"))
			Expect(msg.FirstParams).To(Equal("some"))
			Expect(msg.Params).To(Equal("params"))
		})

		It("handles prefix-less messages", func() {
			msg, err := cipher.Decode("PING some.server\r\n")

			Expect(err).NotTo(HaveOccurred())
			Expect(msg.NickOrSrvname).To(BeZero())
			Expect(msg.Command).To(Equal("PING"))
			Expect(msg.FirstParams).To(Equal("some.server"))
			Expect(msg.Params).To(BeZero())
		})

		It("converts numeric commands to their alphabetic equivalent", func() {
			msg, err := cipher.Decode(":some.server 001 firstparam :some message\r\n")

			Expect(err).NotTo(HaveOccurred())
			Expect(msg.Command).To(Equal("RPL_WELCOME"))
		})

		It("correctly converts PINGs into FirstParam-less messages", func() {
			msg, err := cipher.Decode("PING :some.server\r\n")

			Expect(err).NotTo(HaveOccurred())
			Expect(msg.NickOrSrvname).To(BeZero())
			Expect(msg.Command).To(Equal("PING"))
			Expect(msg.FirstParams).To(BeZero())
			Expect(msg.Params).To(Equal("some.server"))
		})

		It("returns an error when the message is invalid", func() {
			_, err := cipher.Decode("protocols lol")
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("#Encode", func() {
		It("converts a Message into a properly formatted string", func() {
			msg := &message.Message{
				Command:     "SOMECMD",
				FirstParams: "#somechannel",
				Params:      "some message or another",
			}

			str := cipher.Encode(msg)
			Expect(str).To(Equal("SOMECMD #somechannel :some message or another\r\n"))
		})

		It("does not include whitespace or colon if Params is not set", func() {
			msg := &message.Message{
				Command:     "SOMESMALLCMD",
				FirstParams: "#somechannel",
			}

			str := cipher.Encode(msg)
			Expect(str).To(Equal("SOMESMALLCMD #somechannel\r\n"))
		})

		It("does not add extra whitespace if FirstParams is not set", func() {
			msg := &message.Message{
				Command: "SOMESMALLCMD",
				Params:  "some-server",
			}

			str := cipher.Encode(msg)
			Expect(str).To(Equal("SOMESMALLCMD :some-server\r\n"))
		})
	})
})
