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
			app.BeginChatting(fakeRegistrar, fakeSender, cfg)
		})

		It("enrolls with its registrar for login and channeljoin messages", func() {
			Expect(fakeRegistrar.EnrollForEventsCalls).To(Equal(1))
			Expect(fakeRegistrar.EnrollForEventsTypes).To(Equal([]event.Type{event.Login, event.ChannelJoin}))
		})

		It("sets its EventCh with the chan provided by its registrar", func() {
			Expect(app.EventCh).NotTo(Receive())
			eventCh <- &event.Event{}
			Expect(app.EventCh).To(Receive())
		})

		Context("when receiving a login message over the event channel", func() {
			It("sends the appropriate channel-joining params to the provided sender", func() {
				Consistently(fakeSender.SendCalls).Should(Equal(0))
				app.EventCh <- &event.Event{Type: event.Login}

				Eventually(fakeSender.SendCalls).Should(Equal(1))
				Expect(fakeSender.SendCmd).To(Equal("JOIN"))
				Expect(fakeSender.SendFprms).To(Equal("#somechannel"))
				Expect(fakeSender.SendPrms).To(Equal(""))
			})
		})

		Context("before sending its join command", func() {
			It("doesn't respond to channeljoin messages", func() {
				app.EventCh <- &event.Event{Type: event.ChannelJoin}
				Consistently(fakeSender.SendCalls).Should(Equal(0))
			})
		})

		Context("after sending its join command", func() {

			BeforeEach(func() {
				Consistently(fakeSender.SendCalls).Should(Equal(0))
				app.EventCh <- &event.Event{Type: event.Login}
				Eventually(fakeSender.SendCalls).Should(Equal(1))
			})

			Context("when receiving a channeljoin for the correct channel", func() {
				var channeljoinData map[string]string

				BeforeEach(func() {
					channeljoinData = map[string]string{
						"joinedChannel": "somechannel",
					}
					app.EventCh <- &event.Event{event.ChannelJoin, channeljoinData}
				})

				It("announces its arrival", func() {
					Eventually(fakeSender.SendCalls).Should(Equal(2))
					Expect(fakeSender.SendCmd).To(Equal("PRIVMSG"))
					Expect(fakeSender.SendFprms).To(Equal("#somechannel"))
					Expect(fakeSender.SendPrms).To(Equal("COME WITH ME IF YOU WANT TO LIVE."))
				})

				It("stops listening on the channel", func() {
					Eventually(fakeSender.SendCalls).Should(Equal(2))
					Expect(fakeRegistrar.UnsubscribeCalls).To(Equal(1))
					Expect(fakeRegistrar.UnsubscribeChan).To(Equal(app.EventCh))
				})
			})
		})
	})
})
