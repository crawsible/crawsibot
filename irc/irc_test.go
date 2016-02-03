package irc_test

import (
	"bufio"

	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IRC", func() {
	var (
		fakeDialer    *mocks.FakeDialer
		fakeConn      *mocks.FakeConn
		fakeSender    *mocks.FakeSender
		fakeForwarder *mocks.FakeForwarder
		fakePonger    *mocks.FakePonger

		client *irc.IRC
		cfg    *config.Config
	)

	BeforeEach(func() {
		fakeConn = &mocks.FakeConn{}
		fakeDialer = &mocks.FakeDialer{DialReturnConn: fakeConn}
		fakeSender = &mocks.FakeSender{}
		fakeForwarder = &mocks.FakeForwarder{}
		fakePonger = &mocks.FakePonger{}

		client = &irc.IRC{
			Dialer:    fakeDialer,
			Sender:    fakeSender,
			Forwarder: fakeForwarder,
			Ponger:    fakePonger,
		}

		cfg = &config.Config{
			Address:  "some.address:12345",
			Nick:     "some-nick",
			Password: "oauth:key",
		}
	})

	Describe("#Connect", func() {
		It("dials the given address over TCP", func() {
			client.Connect(cfg)

			Expect(fakeDialer.DialCalls).To(Equal(1))
			Expect(fakeDialer.DialNetwork).To(Equal("tcp"))
			Expect(fakeDialer.DialAddress).To(Equal("some.address:12345"))
		})

		It("initiates the Sender with the returned Conn", func() {
			client.Connect(cfg)

			Expect(fakeSender.StartSendingCalls).To(Equal(1))
			Expect(fakeSender.StartSendingWriter).To(Equal(fakeConn))
		})

		It("registers its Ponger for PINGs with the unstarted Forwarder", func() {
			client.Connect(cfg)

			Expect(fakeForwarder.EnrollForPINGCalls).To(Equal(1))
			Expect(fakeForwarder.EnrollForPINGPonger).To(Equal(fakePonger))
			Expect(fakeForwarder.EnrollForPINGPreviousStarts).To(Equal(0))
		})

		It("initiates the Forwarder with a new Reader for the Conn", func() {
			client.Connect(cfg)

			Expect(fakeForwarder.StartForwardingCalls).To(Equal(1))
			fakeReader := bufio.NewReader(fakeConn)
			Expect(fakeForwarder.StartForwardingReader).To(Equal(fakeReader))
		})

		It("sends login messages to the server via the Sender", func() {
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
