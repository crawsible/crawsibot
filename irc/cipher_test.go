package irc_test

import (
	"github.com/crawsible/crawsibot/irc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Cipher", func() {
	Describe("#Decode", func() {
		var (
			cipher *irc.Cipher
			msg    *irc.Message
		)

		BeforeEach(func() {
			msg = cipher.Decode(":some-nick!some-user@some.host SOMECMD some :params\r\n")
		})

		It("parses IRC messages into component parts", func() {
			Expect(msg.NickOrSrvname).To(Equal("some-nick"))
			Expect(msg.Command).To(Equal("SOMECMD"))
			Expect(msg.FirstParam).To(Equal("some"))
			Expect(msg.Params).To(Equal("params"))
		})
	})
})
