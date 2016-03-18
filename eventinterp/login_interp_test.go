package eventinterp_test

import (
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/eventinterp/mocks"
	"github.com/crawsible/crawsibot/irc/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoginInterp", func() {
	var interp *eventinterp.LoginInterp

	Describe("#RegisterForInterp", func() {
		var fakeReceiver1 *mocks.FakeInterpRcvr
		var fakeReceiver2 *mocks.FakeInterpRcvr

		BeforeEach(func() {
			fakeReceiver1 = &mocks.FakeInterpRcvr{}
			fakeReceiver2 = &mocks.FakeInterpRcvr{}
			interp = &eventinterp.LoginInterp{}
		})

		It("adds the provided receiver to its list of LoginRcvrs", func() {
			interp.RegisterForInterp(fakeReceiver1)
			Expect(interp.LoginRcvrs).To(Equal([]eventinterp.LoginRcvr{fakeReceiver1}))
			interp.RegisterForInterp(fakeReceiver2)
			Expect(interp.LoginRcvrs).To(Equal([]eventinterp.LoginRcvr{
				fakeReceiver1,
				fakeReceiver2,
			}))
		})
	})

	Describe("#BeginInterpreting", func() {
		var (
			eventCh      chan *models.Message
			fakeEnroller *mocks.FakeEnroller
		)

		BeforeEach(func() {
			eventCh = make(chan *models.Message, 1)
			fakeEnroller = &mocks.FakeEnroller{EnrollForMsgsReturnChan: eventCh}
			interp = &eventinterp.LoginInterp{}
		})

		JustBeforeEach(func() {
			interp.BeginInterpreting(fakeEnroller)
		})

		It("sets its EventCh with the chan provided by its enroller", func() {
			Expect(fakeEnroller.EnrollForMsgsCalls).To(Equal(1))
			Expect(fakeEnroller.EnrollForMsgsCmd).To(Equal("RPL_ENDOFMOTD"))

			Expect(interp.EventCh).NotTo(Receive())
			eventCh <- &models.Message{}
			Expect(interp.EventCh).To(Receive())
		})

		Context("when receiving a message over the event channel", func() {
			var (
				fakeReceiver1 *mocks.FakeInterpRcvr
				fakeReceiver2 *mocks.FakeInterpRcvr
			)

			BeforeEach(func() {
				fakeReceiver1 = &mocks.FakeInterpRcvr{}
				fakeReceiver2 = &mocks.FakeInterpRcvr{}

				interp = &eventinterp.LoginInterp{
					LoginRcvrs: []eventinterp.LoginRcvr{fakeReceiver1, fakeReceiver2},
				}
			})

			It("invokes the 'LoggedIn' method of its registered InterpRcvrs", func() {
				Eventually(fakeReceiver1.LoggedInCalls).ShouldNot(Equal(1))
				Eventually(fakeReceiver2.LoggedInCalls).ShouldNot(Equal(1))
				interp.EventCh <- &models.Message{}
				Eventually(fakeReceiver1.LoggedInCalls).Should(Equal(1))
				Eventually(fakeReceiver2.LoggedInCalls).Should(Equal(1))
			})
		})
	})
})
