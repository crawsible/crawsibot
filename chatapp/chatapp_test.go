package chatapp_test

import (
	"github.com/crawsible/crawsibot/chatapp"
	"github.com/crawsible/crawsibot/chatapp/mocks"

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
			fakeApp        *mocks.FakeApp
			responder      *chatapp.ChatApp
		)

		BeforeEach(func() {
			fakeController = &mocks.FakeController{}
			fakeApp = &mocks.FakeApp{}

			responder = &chatapp.ChatApp{
				JoinChannelApp: fakeApp,
			}
		})

		It("tells its JoinChannelApp to begin chatting", func() {
			responder.BeginChatting(fakeController)
			Expect(fakeApp.BeginChattingCalls).To(Equal(1))
			Expect(fakeApp.BeginChattingRegistrar).To(Equal(fakeController))
		})
	})
})
