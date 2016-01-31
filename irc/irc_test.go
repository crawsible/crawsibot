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
		ircClient  *irc.IRC
		cfg        *config.Config
	)

	BeforeEach(func() {
		fakeDialer = &mocks.FakeDialer{}

		ircClient = &irc.IRC{
			Dialer: fakeDialer,
		}
		cfg = &config.Config{
			Address: "some.address:12345",
		}
	})

	Describe("#Connect", func() {
		It("should dial the given address over tcp", func() {
			ircClient.Connect(cfg)
			Expect(fakeDialer.DialCalls).To(Equal(1))

			Expect(fakeDialer.DialNetwork).To(Equal("tcp"))
			Expect(fakeDialer.DialAddress).To(Equal("some.address:12345"))
		})
	})
})
