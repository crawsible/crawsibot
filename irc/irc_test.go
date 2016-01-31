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
		fakeCipher *mocks.FakeCipher

		ircClient *irc.IRC
		cfg       *config.Config
	)

	BeforeEach(func() {
		fakeDialer = &mocks.FakeDialer{}
		fakeConn = &mocks.FakeConn{}
		fakeCipher = &mocks.FakeCipher{}

		fakeDialer.DialReturnConn = fakeConn

		ircClient = &irc.IRC{
			Dialer: fakeDialer,
			Cipher: fakeCipher,
		}

		cfg = &config.Config{
			Address:  "some.address:12345",
			Nick:     "some-nick",
			Password: "oauth:key",
		}
	})

	Describe("#Connect", func() {
		It("should dial the given address over tcp", func() {
			ircClient.Connect(cfg)

			Expect(fakeDialer.DialCalls).To(Equal(1))
			Expect(fakeDialer.DialNetwork).To(Equal("tcp"))
			Expect(fakeDialer.DialAddress).To(Equal("some.address:12345"))
		})

		It("should use the cipher to generate validation message strings", func() {
			ircClient.Connect(cfg)

			Expect(fakeCipher.EncodeCalls).To(Equal(2))
			msg1 := fakeCipher.EncodeMessages[0]
			msg2 := fakeCipher.EncodeMessages[1]

			Expect(msg1.Command).To(Equal("PASS"))
			Expect(msg1.FirstParams).To(Equal("oauth:key"))
			Expect(msg2.Command).To(Equal("NICK"))
			Expect(msg2.FirstParams).To(Equal("some-nick"))
		})

		It("should validate against the conn with cipher-encoded messages", func() {
			fakeCipher.EncodeReturns = append(
				fakeCipher.EncodeReturns,
				"some-encoded-string1",
				"some-encoded-string2",
			)
			ircClient.Connect(cfg)

			Expect(fakeConn.WriteCalls).To(Equal(1))
			authMsg := []byte("some-encoded-string1some-encoded-string2")
			Expect(fakeConn.WriteMessage).To(Equal(authMsg))
		})
	})
})
