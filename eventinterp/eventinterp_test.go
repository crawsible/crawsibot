package eventinterp_test

import (
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/eventinterp/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventInterp", func() {
	Describe(".New", func() {
		It("returns an EventInterp with default dependencies injected", func() {
			c := eventinterp.New()
			Expect(c.LoginInterp).To(Equal(&eventinterp.LoginInterp{}))
		})
	})

	Describe("#BeginInterpreting", func() {
		var (
			fakeClient *mocks.FakeClient
			fakeInterp *mocks.FakeInterp
			controller *eventinterp.EventInterp
		)

		BeforeEach(func() {
			fakeClient = &mocks.FakeClient{}
			fakeInterp = &mocks.FakeInterp{}

			controller = &eventinterp.EventInterp{
				LoginInterp: fakeInterp,
			}
		})

		It("tells its LoginInterp to begin interpreting", func() {
			controller.BeginInterpreting(fakeClient)
			Expect(fakeInterp.BeginInterpretingCalls).To(Equal(1))
			Expect(fakeInterp.BeginInterpretingClient).To(Equal(fakeClient))
		})
	})
})
