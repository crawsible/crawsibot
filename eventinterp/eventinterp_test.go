package eventinterp_test

import (
	"github.com/crawsible/crawsibot/eventinterp"

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
})
