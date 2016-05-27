package eventinterp_test

import (
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/eventinterp/event"
	"github.com/crawsible/crawsibot/eventinterp/mocks"
	"github.com/crawsible/crawsibot/irc/message"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ChannelJoinInterp", func() {
	var interp *eventinterp.ChannelJoinInterp

	Describe("#RegisterForInterp", func() {
		var eventCh1 chan *event.Event
		var eventCh2 chan *event.Event

		BeforeEach(func() {
			eventCh1 = make(chan *event.Event, 1)
			eventCh2 = make(chan *event.Event, 1)
			interp = &eventinterp.ChannelJoinInterp{}
		})

		It("adds the provided receiver to its list of EventChs", func() {
			interp.RegisterForInterp(eventCh1)
			interp.RegisterForInterp(eventCh2)

			Expect(eventCh1).NotTo(Receive())
			Expect(eventCh2).NotTo(Receive())

			interp.EventChs[0] <- &event.Event{}
			Expect(eventCh1).To(Receive())
			Expect(eventCh2).NotTo(Receive())

			interp.EventChs[1] <- &event.Event{}
			Expect(eventCh1).NotTo(Receive())
			Expect(eventCh2).To(Receive())
		})
	})

	Describe("#BeginInterpreting", func() {
		var (
			msgCh        chan *message.Message
			fakeEnroller *mocks.FakeEnroller
		)

		BeforeEach(func() {
			msgCh = make(chan *message.Message, 1)
			fakeEnroller = &mocks.FakeEnroller{EnrollForMsgsReturnChan: msgCh}
			interp = &eventinterp.ChannelJoinInterp{}
		})

		JustBeforeEach(func() {
			interp.BeginInterpreting(fakeEnroller)
		})

		It("sets its MsgCh with the chan provided by its enroller", func() {
			Expect(fakeEnroller.EnrollForMsgsCalls).To(Equal(1))
			Expect(fakeEnroller.EnrollForMsgsCmd).To(Equal("RPL_ENDOFNAMES"))

			Expect(interp.MsgCh).NotTo(Receive())
			msgCh <- &message.Message{}
			Expect(interp.MsgCh).To(Receive())
		})

		Context("when receiving a message over the event channel", func() {
			var eventCh1 chan *event.Event
			var eventCh2 chan *event.Event

			BeforeEach(func() {
				eventCh1 = make(chan *event.Event, 1)
				eventCh2 = make(chan *event.Event, 1)
				eventChs := []chan *event.Event{eventCh1, eventCh2}

				interp.EventChs = eventChs
			})

			It("sends an event to its eventChs", func() {
				Consistently(eventCh1).ShouldNot(Receive())
				Consistently(eventCh2).ShouldNot(Receive())

				msgCh <- &message.Message{}
				Eventually(eventCh1).Should(Receive())
				Eventually(eventCh2).Should(Receive())
			})

			It("includes the channel as Event data", func() {
				msgCh <- &message.Message{
					FirstParams: "someusername #somechannel",
				}

				expectedEvent := &event.Event{
					Type: event.ChannelJoin,
					Data: map[string]string{
						"joinedChannel": "somechannel",
					},
				}

				Eventually(eventCh1).Should(Receive(Equal(expectedEvent)))
			})
		})
	})

	Describe("#Unsubscribe", func() {
		var eventCh1 chan *event.Event
		var eventCh2 chan *event.Event
		var eventCh3 chan *event.Event

		BeforeEach(func() {
			eventCh1 = make(chan *event.Event, 1)
			eventCh2 = make(chan *event.Event, 1)
			eventCh3 = make(chan *event.Event, 1)
			eventChs := []chan *event.Event{eventCh1, eventCh2, eventCh3}

			interp = &eventinterp.ChannelJoinInterp{}
			interp.EventChs = eventChs
		})

		It("removes a registered channel from the EventChs slice", func() {
			interp.Unsubscribe(eventCh2)
			Expect(interp.EventChs).To(HaveLen(2))

			Expect(eventCh1).NotTo(Receive())
			Expect(eventCh2).NotTo(Receive())
			Expect(eventCh3).NotTo(Receive())

			interp.EventChs[0] <- &event.Event{}
			Expect(eventCh1).To(Receive())
			interp.EventChs[1] <- &event.Event{}
			Expect(eventCh3).To(Receive())
		})
	})
})
