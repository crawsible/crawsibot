package eventinterp_test

import (
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/eventinterp/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoginAnalyst", func() {
	var analyst *eventinterp.LoginAnalyst

	Describe("#BeginInterpreting", func() {
		var fakeClient *mocks.FakeClient

		BeforeEach(func() {
			fakeClient = &mocks.FakeClient{}
			analyst = &eventinterp.LoginAnalyst{}
		})

		JustBeforeEach(func() {
			analyst.BeginInterpreting(fakeClient)
		})

		It("instantiates its event channel with a buffer of 1", func() {
			Expect(analyst.EventCh).NotTo(BeNil())
			Expect(cap(analyst.EventCh)).To(Equal(1))
		})

		It("registers itself for RPL_ENDOFMOTD messages with the client", func() {
			Expect(fakeClient.EnrollForMsgsCalls).To(Equal(1))
			Expect(fakeClient.EnrollForMsgsRcvr).To(Equal(analyst))
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

				analyst = &eventinterp.LoginAnalyst{
					LoginRcvrs: []eventinterp.LoginRcvr{fakeReceiver1, fakeReceiver2},
					EventCh:    make(chan struct{}, 1),
				}
			})

			It("invokes the 'LoggedIn' method of its registered InterpRcvrs", func() {
				Eventually(fakeReceiver1.LoggedInCalls).ShouldNot(Equal(1))
				Eventually(fakeReceiver2.LoggedInCalls).ShouldNot(Equal(1))
				analyst.EventCh <- struct{}{}
				Eventually(fakeReceiver1.LoggedInCalls).Should(Equal(1))
				Eventually(fakeReceiver2.LoggedInCalls).Should(Equal(1))
			})
		})
	})

	Describe("#RcvMsg", func() {
		BeforeEach(func() {
			analyst = &eventinterp.LoginAnalyst{EventCh: make(chan struct{}, 1)}
		})

		It("sends on the analyst's EventCh", func() {
			analyst.RcvMsg("", "", "")
			Eventually(analyst.EventCh).Should(Receive())
		})
	})
})
