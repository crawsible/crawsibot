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
			Expect(c.LoginAnalyst).To(Equal(&eventinterp.LoginAnalyst{}))
		})
	})

	Describe("#BeginInterpreting", func() {
		var (
			fakeClient  *mocks.FakeClient
			fakeAnalyst *mocks.FakeAnalyst
			controller  *eventinterp.EventInterp
		)

		BeforeEach(func() {
			fakeClient = &mocks.FakeClient{}
			fakeAnalyst = &mocks.FakeAnalyst{}

			controller = &eventinterp.EventInterp{
				LoginAnalyst: fakeAnalyst,
			}
		})

		It("tells its LoginAnalyst to begin interpreting", func() {
			controller.BeginInterpreting(fakeClient)
			Expect(fakeAnalyst.BeginInterpretingCalls).To(Equal(1))
			Expect(fakeAnalyst.BeginInterpretingClient).To(Equal(fakeClient))
		})
	})
})
