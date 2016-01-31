package irc_test

import (
	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IRC", func() {
	var (
		fakeDialer *mocks.FakeDialer
		fakeConn   *mocks.FakeConn

		ircClient *irc.IRC
		cfg       *config.Config
	)

	BeforeEach(func() {
		fakeDialer = &mocks.FakeDialer{}
		fakeConn = &mocks.FakeConn{}

		fakeDialer.DialReturnConn = fakeConn

		ircClient = &irc.IRC{
			Dialer: fakeDialer,
		}
		cfg = &config.Config{
			Address:  "some.address:12345",
			Nick:     "some-nick",
			Password: "oauth:key",
		}
	})

	Describe("#Connect", func() {
		BeforeEach(func() {
			ircClient.Connect(cfg)
		})

		It("should dial the given address over tcp", func() {
			Expect(fakeDialer.DialCalls).To(Equal(1))
			Expect(fakeDialer.DialNetwork).To(Equal("tcp"))
			Expect(fakeDialer.DialAddress).To(Equal("some.address:12345"))
		})

		It("should validate against the server", func() {
			Expect(fakeConn.WriteCalls).To(Equal(1))

			authMsg := []byte("PASS oauth:key\r\nNICK some-nick\r\n")
			Expect(fakeConn.WriteMessage).To(Equal(authMsg))
		})
	})
})
