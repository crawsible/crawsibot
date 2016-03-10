package eventinterp_test

import (
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/eventinterp/mocks"

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
		var fakeClient *mocks.FakeClient

		BeforeEach(func() {
			fakeClient = &mocks.FakeClient{}
			interp = &eventinterp.LoginInterp{}
		})

		JustBeforeEach(func() {
			interp.BeginInterpreting(fakeClient)
		})

		It("instantiates its event channel with a buffer of 1", func() {
			Expect(interp.EventCh).NotTo(BeNil())
			Expect(cap(interp.EventCh)).To(Equal(1))
		})

		It("registers itself for RPL_ENDOFMOTD messages with the client", func() {
			Expect(fakeClient.EnrollForMsgsCalls).To(Equal(1))
			Expect(fakeClient.EnrollForMsgsRcvr).To(Equal(interp))
			Expect(fakeClient.EnrollForMsgsCmd).To(Equal("RPL_ENDOFMOTD"))
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
					EventCh:    make(chan struct{}, 1),
				}
			})

			It("invokes the 'LoggedIn' method of its registered InterpRcvrs", func() {
				Eventually(fakeReceiver1.LoggedInCalls).ShouldNot(Equal(1))
				Eventually(fakeReceiver2.LoggedInCalls).ShouldNot(Equal(1))
				interp.EventCh <- struct{}{}
				Eventually(fakeReceiver1.LoggedInCalls).Should(Equal(1))
				Eventually(fakeReceiver2.LoggedInCalls).Should(Equal(1))
			})
		})
	})

	Describe("#RcvMsg", func() {
		BeforeEach(func() {
			interp = &eventinterp.LoginInterp{EventCh: make(chan struct{}, 1)}
		})

		It("sends on the interp's EventCh", func() {
			interp.RcvMsg("", "", "")
			Eventually(interp.EventCh).Should(Receive())
		})
	})
})
