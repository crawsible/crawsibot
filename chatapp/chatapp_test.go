package chatapp_test

import (
	"github.com/crawsible/crawsibot/chatapp"
	"github.com/crawsible/crawsibot/chatapp/mocks"
	"github.com/crawsible/crawsibot/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ChatApp", func() {
	Describe(".New", func() {
		It("returns a ChatApp with default dependencies injected", func() {
			r := chatapp.New()
			Expect(r.JoinChannelApp).To(Equal(&chatapp.JoinChannelApp{}))
		})
	})

	Describe("#BeginChatting", func() {
		var (
			fakeController *mocks.FakeController
			fakeSender     *mocks.FakeSender
			fakeApp        *mocks.FakeApp
			cfg            *config.Config

			responder *chatapp.ChatApp
		)

		BeforeEach(func() {
			fakeController = &mocks.FakeController{}
			fakeSender = &mocks.FakeSender{}
			fakeApp = &mocks.FakeApp{}
			cfg = &config.Config{Channel: "somechannel"}

			responder = &chatapp.ChatApp{
				JoinChannelApp: fakeApp,
			}
			responder.BeginChatting(fakeController, fakeSender, cfg)
		})

		It("tells its JoinChannelApp to begin chatting", func() {
			Expect(fakeApp.BeginChattingCalls).To(Equal(1))
			Expect(fakeApp.BeginChattingRegistrar).To(Equal(fakeController))
			Expect(fakeApp.BeginChattingSender).To(Equal(fakeSender))
			Expect(fakeApp.BeginChattingCfg).To(Equal(cfg))
		})
	})
})
