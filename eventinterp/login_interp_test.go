package eventinterp_test

import (
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/eventinterp/event"
	"github.com/crawsible/crawsibot/eventinterp/mocks"
	"github.com/crawsible/crawsibot/irc/message"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoginInterp", func() {
	var interp *eventinterp.LoginInterp

	Describe("#RegisterForInterp", func() {
		var eventCh1 chan *event.Event
		var eventCh2 chan *event.Event

		BeforeEach(func() {
			eventCh1 = make(chan *event.Event, 1)
			eventCh2 = make(chan *event.Event, 1)
			interp = &eventinterp.LoginInterp{}
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
			interp = &eventinterp.LoginInterp{}
		})

		JustBeforeEach(func() {
			interp.BeginInterpreting(fakeEnroller)
		})

		It("sets its MsgCh with the chan provided by its enroller", func() {
			Expect(fakeEnroller.EnrollForMsgsCalls).To(Equal(1))
			Expect(fakeEnroller.EnrollForMsgsCmd).To(Equal("RPL_ENDOFMOTD"))

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

				interp = &eventinterp.LoginInterp{
					EventChs: []chan *event.Event{eventCh1, eventCh2},
				}
			})

			It("sends a login Event to its eventChs", func() {
				Expect(eventCh1).NotTo(Receive())
				Expect(eventCh2).NotTo(Receive())

				msgCh <- &message.Message{}
				Eventually(eventCh1).Should(Receive(Equal(&event.Event{Type: event.Login})))
				Eventually(eventCh2).Should(Receive(Equal(&event.Event{Type: event.Login})))
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

			interp = &eventinterp.LoginInterp{}
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
