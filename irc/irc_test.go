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
		fakeSender *mocks.FakeSender

		client *irc.IRC
		cfg    *config.Config
	)

	BeforeEach(func() {
		fakeDialer = &mocks.FakeDialer{}
		fakeConn = &mocks.FakeConn{}
		fakeDialer.DialReturnConn = fakeConn
		fakeSender = &mocks.FakeSender{}

		client = &irc.IRC{
			Dialer: fakeDialer,
			Sender: fakeSender,
		}

		cfg = &config.Config{
			Address:  "some.address:12345",
			Nick:     "some-nick",
			Password: "oauth:key",
		}
	})

	Describe("#Connect", func() {
		It("dials the given address over tcp", func() {
			client.Connect(cfg)

			Expect(fakeDialer.DialCalls).To(Equal(1))
			Expect(fakeDialer.DialNetwork).To(Equal("tcp"))
			Expect(fakeDialer.DialAddress).To(Equal("some.address:12345"))
		})

		It("initiates the sender with the returned conn", func() {
			client.Connect(cfg)

			Expect(fakeSender.StartSendingCalls).To(Equal(1))
			Expect(fakeSender.StartSendingConn).To(Equal(fakeConn))
		})

		It("sends login messages to the server via the sender", func() {
			msgCh := make(chan *irc.Message)
			defer close(msgCh)
			fakeSender.ReturnCh = msgCh

			client.Connect(cfg)

			Eventually(fakeSender.ReceivedOverChan).Should(HaveLen(3))

			passMsg := &irc.Message{Command: "PASS", FirstParams: cfg.Password}
			nickMsg := &irc.Message{Command: "NICK", FirstParams: cfg.Nick}
			capMsg := &irc.Message{
				Command:     "CAP",
				FirstParams: "REQ",
				Params:      "twitch.tv/membership",
			}
			Expect(fakeSender.ReceivedOverChan()[0]).To(Equal(passMsg))
			Expect(fakeSender.ReceivedOverChan()[1]).To(Equal(nickMsg))
			Expect(fakeSender.ReceivedOverChan()[2]).To(Equal(capMsg))
		})
	})
})
