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

	Context("with EventInterp instances", func() {
		var (
			fakeInterp *mocks.FakeInterp
			controller *eventinterp.EventInterp
		)

		BeforeEach(func() {
			fakeInterp = &mocks.FakeInterp{}
			controller = &eventinterp.EventInterp{
				LoginInterp: fakeInterp,
			}
		})

		Describe("#BeginInterpreting", func() {
			var fakeEnroller *mocks.FakeEnroller

			BeforeEach(func() {
				fakeEnroller = &mocks.FakeEnroller{}
				controller.BeginInterpreting(fakeEnroller)
			})

			It("tells its LoginInterp to begin interpreting", func() {
				Expect(fakeInterp.BeginInterpretingCalls).To(Equal(1))
				Expect(fakeInterp.BeginInterpretingEnroller).To(Equal(fakeEnroller))
			})
		})

		Describe("#RegisterForLogin", func() {
			var fakeRcvr *mocks.FakeInterpRcvr

			BeforeEach(func() {
				fakeRcvr = &mocks.FakeInterpRcvr{}
				controller.RegisterForLogin(fakeRcvr)
			})

			It("registers the provided LoginRcvr with its LoginInterp", func() {
				Expect(fakeInterp.RegisterForInterpCalls).To(Equal(1))
				Expect(fakeInterp.RegisterForInterpRcvr).To(Equal(fakeRcvr))
			})
		})
	})
})
