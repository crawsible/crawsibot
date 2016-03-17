package irc_test

import (
	"bufio"
	"net"

	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IRC", func() {
	Describe(".New", func() {
		It("returns an IRC client with default dependencies injected", func() {
			c := irc.New()

			Expect(c.Dialer).To(Equal(&net.Dialer{}))
			Expect(c.Cipher).To(Equal(&irc.MessageCipher{}))
			Expect(c.Sender).To(Equal(&irc.MessageSender{}))
			Expect(c.Forwarder).To(Equal(&irc.MessageForwarder{MsgRcvrs: make(map[string][]irc.MsgRcvr)}))
			Expect(c.Ponger).To(Equal(&irc.ServerPonger{}))
		})
	})

	Context("with an IRC instance", func() {
		var (
			fakeConn      *mocks.FakeConn
			fakeDialer    *mocks.FakeDialer
			fakeCipher    *mocks.FakeCipher
			fakeSender    *mocks.FakeSender
			fakeForwarder *mocks.FakeForwarder
			fakePonger    *mocks.FakePonger

			cfg    *config.Config
			client *irc.IRC
		)

		BeforeEach(func() {
			fakeConn = &mocks.FakeConn{}
			fakeDialer = &mocks.FakeDialer{DialReturnConn: fakeConn}
			fakeCipher = &mocks.FakeCipher{}
			fakeSender = &mocks.FakeSender{}
			fakeForwarder = &mocks.FakeForwarder{}
			fakePonger = &mocks.FakePonger{FakeForwarder: fakeForwarder}

			client = &irc.IRC{
				Dialer:    fakeDialer,
				Cipher:    fakeCipher,
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

			It("starts the Sender with its cipher and returned Conn", func() {
				client.Connect(cfg)

				Expect(fakeSender.StartSendingCalls).To(Equal(1))
				Expect(fakeSender.StartSendingWriter).To(Equal(fakeConn))
				Expect(fakeSender.StartSendingEcdr).To(Equal(fakeCipher))
			})

			It("starts its Ponger with itself as the API argument", func() {
				client.Connect(cfg)

				Expect(fakePonger.StartPongingCalls).To(Equal(1))
				Expect(fakePonger.StartPongingMsgr).To(Equal(client))
				Expect(fakePonger.ForwarderHadStarted).To(BeFalse())
			})

			It("starts the Forwarder with the cipher and a Conn Reader", func() {
				client.Connect(cfg)

				Expect(fakeForwarder.StartForwardingCalls).To(Equal(1))
				fakeReader := bufio.NewReader(fakeConn)
				Expect(fakeForwarder.StartForwardingReader).To(Equal(fakeReader))
				Expect(fakeForwarder.StartForwardingDcdr).To(Equal(fakeCipher))
			})

			It("sends login messages to the server via the Sender", func() {
				client.Connect(cfg)

				Expect(fakeSender.SendCalls).Should(Equal(3))
				passArgs := []string{"PASS", cfg.Password, ""}
				nickArgs := []string{"NICK", cfg.Nick, ""}
				capArgs := []string{"CAP", "REQ", "twitch.tv/membership"}

				Expect(fakeSender.SendArgs[0]).To(Equal(passArgs))
				Expect(fakeSender.SendArgs[1]).To(Equal(nickArgs))
				Expect(fakeSender.SendArgs[2]).To(Equal(capArgs))
			})
		})

		Describe("#EnrollForMsgs", func() {
			FIt("calls through to its Forwarder", func() {
				rcvCh := make(chan map[string]string)
				client.EnrollForMsgs(rcvCh, "SOMECMD")

				Expect(fakeForwarder.EnrollForMsgCalls).To(Equal(1))
				Expect(fakeForwarder.EnrollForMsgChan).To(Equal(rcvCh))
				Expect(fakeForwarder.EnrollForMsgCmd).To(Equal("SOMECMD"))
			})
		})

		Describe("#Send", func() {
			It("calls through to its sender", func() {
				client.Send("some-cmd", "some-fprms", "some-prms")

				Expect(fakeSender.SendCalls).To(Equal(1))
				expectedArgs := []string{"some-cmd", "some-fprms", "some-prms"}
				Expect(fakeSender.SendArgs[0]).To(Equal(expectedArgs))
			})
		})
	})
})
