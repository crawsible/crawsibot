package eventinterp_test

import (
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/eventinterp/event"
	"github.com/crawsible/crawsibot/eventinterp/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventInterp", func() {
	Describe(".New", func() {
		It("returns an EventInterp with default dependencies injected", func() {
			c := eventinterp.New()
			Expect(c.Interps).To(Equal(
				map[event.Type]eventinterp.Interp{
					event.Login: &eventinterp.LoginInterp{},
				},
			))
		})
	})

	Context("with EventInterp instances", func() {
		var (
			fakeInterp    *mocks.FakeInterp
			anotherInterp *mocks.FakeInterp
			controller    *eventinterp.EventInterp
		)

		BeforeEach(func() {
			fakeInterp = &mocks.FakeInterp{}
			anotherInterp = &mocks.FakeInterp{}
			controller = &eventinterp.EventInterp{
				Interps: map[event.Type]eventinterp.Interp{
					event.Login:   fakeInterp,
					event.Unknown: anotherInterp,
				},
			}
		})

		Describe("#BeginInterpreting", func() {
			var fakeEnroller *mocks.FakeEnroller

			BeforeEach(func() {
				fakeEnroller = &mocks.FakeEnroller{}
				controller.BeginInterpreting(fakeEnroller)
			})

			It("tells each of its Interps to begin interpreting", func() {
				Expect(fakeInterp.BeginInterpretingCalls).To(Equal(1))
				Expect(fakeInterp.BeginInterpretingEnroller).To(Equal(fakeEnroller))
				Expect(anotherInterp.BeginInterpretingCalls).To(Equal(1))
				Expect(anotherInterp.BeginInterpretingEnroller).To(Equal(fakeEnroller))
			})
		})

		Describe("#EnrollForEvents", func() {
			It("generates and returns a single-buffered event chan", func() {
				ch := controller.EnrollForEvents(event.Login)

				var eventCh chan *event.Event
				Expect(ch).To(BeAssignableToTypeOf(eventCh))
				Expect(cap(ch)).To(Equal(1))
			})

			It("enrolls the channel with the interps for the provided event Types", func() {
				ch := controller.EnrollForEvents(event.Login)

				Expect(fakeInterp.RegisterForInterpCalls).To(Equal(1))

				Expect(ch).NotTo(Receive())
				fakeInterp.RegisterForInterpChan <- &event.Event{}
				Expect(ch).To(Receive())
			})

			It("will enroll a single channel for multiple Types", func() {
				ch := controller.EnrollForEvents(event.Login, event.Unknown)

				Expect(fakeInterp.RegisterForInterpCalls).To(Equal(1))
				Expect(anotherInterp.RegisterForInterpCalls).To(Equal(1))

				Expect(ch).NotTo(Receive())
				fakeInterp.RegisterForInterpChan <- &event.Event{}
				Expect(ch).To(Receive())
				anotherInterp.RegisterForInterpChan <- &event.Event{}
				Expect(ch).To(Receive())
			})
		})
	})
})
