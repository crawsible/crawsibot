package irc_test

import (
	"github.com/crawsible/crawsibot/irc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cipher", func() {
	Describe("#Decode", func() {
		var cipher *irc.Cipher

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
			msg, err := cipher.Decode(":some.server 001 first-param :Welcome message\r\n")

			Expect(err).NotTo(HaveOccurred())
			Expect(msg.Command).To(Equal("RPL_WELCOME"))
		})

		It("returns an error when the message is invalid", func() {
			_, err := cipher.Decode("protocols lol")
			Expect(err).To(HaveOccurred())
		})
	})
})
