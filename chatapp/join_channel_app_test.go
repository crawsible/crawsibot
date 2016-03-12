package chatapp_test

import (
	"github.com/crawsible/crawsibot/chatapp"
	"github.com/crawsible/crawsibot/chatapp/mocks"
	"github.com/crawsible/crawsibot/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JoinChannelApp", func() {
	var app *chatapp.JoinChannelApp

	Describe("#BeginChatting", func() {
		var (
			fakeRegistrar *mocks.FakeRegistrar
			fakeSender    *mocks.FakeSender
			cfg           *config.Config
		)

		BeforeEach(func() {
			fakeRegistrar = &mocks.FakeRegistrar{}
			fakeSender = &mocks.FakeSender{}
			cfg = &config.Config{Channel: "somechannel"}

			app = &chatapp.JoinChannelApp{}
		})

		JustBeforeEach(func() {
			app.BeginChatting(fakeRegistrar, fakeSender, cfg)
		})

		It("instantiates its event channel with a buffer of 1", func() {
			Expect(app.EventCh).NotTo(BeNil())
			Expect(cap(app.EventCh)).To(Equal(1))
		})

		It("registers itself for Login messages with the provided registrar", func() {
			Expect(fakeRegistrar.RegisterForLoginCalls).To(Equal(1))
			Expect(fakeRegistrar.RegisterForLoginRcvr).To(Equal(app))
		})

		Context("when receiving a message over the event channel", func() {
			It("sends the appropriate channel-joining params to the provided sender", func() {
				Eventually(fakeSender.SendCalls).ShouldNot(Equal(1))
				app.EventCh <- struct{}{}
				Eventually(fakeSender.SendCalls).Should(Equal(1))
				Expect(fakeSender.SendCmd).To(Equal("JOIN"))
				Expect(fakeSender.SendFprms).To(Equal("#somechannel"))
				Expect(fakeSender.SendPrms).To(Equal(""))
			})
		})
	})

	Describe("#LoggedIn", func() {
		BeforeEach(func() {
			app = &chatapp.JoinChannelApp{EventCh: make(chan struct{}, 1)}
		})

		It("sends on the app's EventCh", func() {
			app.LoggedIn()
			Eventually(app.EventCh).Should(Receive())
		})
	})
})
