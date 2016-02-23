package eventinterp_test

import (
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/eventinterp/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoginAnalyst", func() {
	Describe("#BeginInterpreting", func() {
		var (
			fakeClient *mocks.FakeClient
			analyst    *eventinterp.LoginAnalyst
		)

		BeforeEach(func() {
			fakeClient = &mocks.FakeClient{}
			analyst = &eventinterp.LoginAnalyst{}
		})

		It("registers itself for RPL_ENDOFMOTD messages with the client", func() {
			analyst.BeginInterpreting(fakeClient)

			Expect(fakeClient.EnrollForMsgsCalls).To(Equal(1))
			Expect(fakeClient.EnrollForMsgsRcvr).To(Equal(analyst))
			Expect(fakeClient.EnrollForMsgsCmd).To(Equal("RPL_ENDOFMOTD"))
		})
	})
})
