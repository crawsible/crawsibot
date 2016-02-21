package eventinterp_test

import (
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/eventinterp/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventInterp", func() {
	Describe(".New", func() {
		It("returns an eventinterp with default dependencies injected", func() {
			c := eventinterp.New()
			Expect(c).To(Equal(&eventinterp.EventInterp{}))
		})
	})

	Describe("#BeginInterpreting", func() {
		var (
			fakeClient *mocks.FakeClient
			controller *eventinterp.EventInterp
		)

		BeforeEach(func() {
			fakeClient = &mocks.FakeClient{}
			controller = &eventinterp.EventInterp{}
		})

		It("registers itself for RPL_ENDOFMOTD messages with the client", func() {
			controller.BeginInterpreting(fakeClient)
			Expect(fakeClient.EnrollForMsgsCalls).To(Equal(1))
			Expect(fakeClient.EnrollForMsgsRcvr).To(Equal(controller))
			Expect(fakeClient.EnrollForMsgsCmd).To(Equal("RPL_ENDOFMOTD"))
		})
	})
})
