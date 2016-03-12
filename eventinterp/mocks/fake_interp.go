package mocks

import "github.com/crawsible/crawsibot/eventinterp"

type FakeInterp struct {
	BeginInterpretingCalls  int
	BeginInterpretingClient eventinterp.Enroller
}

func (a *FakeInterp) BeginInterpreting(enlr eventinterp.Enroller) {
	a.BeginInterpretingCalls += 1
	a.BeginInterpretingClient = enlr
}
