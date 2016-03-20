package chatapp_test

import (
	"github.com/crawsible/crawsibot/chatapp"
	"github.com/crawsible/crawsibot/chatapp/mocks"
	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/eventinterp/event"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JoinChannelApp", func() {
	var app *chatapp.JoinChannelApp

	Describe("#BeginChatting", func() {
		var (
			fakeRegistrar *mocks.FakeRegistrar
			eventCh       chan *event.Event
			fakeSender    *mocks.FakeSender
			cfg           *config.Config
		)

		BeforeEach(func() {
			eventCh = make(chan *event.Event, 1)
			fakeRegistrar = &mocks.FakeRegistrar{EnrollForEventsReturnChan: eventCh}
			fakeSender = &mocks.FakeSender{}
			cfg = &config.Config{Channel: "somechannel"}

			app = &chatapp.JoinChannelApp{}
		})

		JustBeforeEach(func() {
			app.BeginChatting(fakeRegistrar, fakeSender, cfg)
		})

		It("enrolls with its registrar for login messages", func() {
			Expect(fakeRegistrar.EnrollForEventsCalls).To(Equal(1))
			Expect(fakeRegistrar.EnrollForEventsTypes).To(Equal([]event.Type{event.Login}))
		})

		It("sets its EventCh with the chan provided by its registrar", func() {
			Expect(app.EventCh).NotTo(Receive())
			eventCh <- &event.Event{}
			Expect(app.EventCh).To(Receive())
		})

		Context("when receiving a message over the event channel", func() {
			It("sends the appropriate channel-joining params to the provided sender", func() {
				Eventually(fakeSender.SendCalls).ShouldNot(Equal(1))
				app.EventCh <- &event.Event{}
				Eventually(fakeSender.SendCalls).Should(Equal(1))
				Expect(fakeSender.SendCmd).To(Equal("JOIN"))
				Expect(fakeSender.SendFprms).To(Equal("#somechannel"))
				Expect(fakeSender.SendPrms).To(Equal(""))
			})
		})
	})
})
