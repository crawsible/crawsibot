package mocks

import "github.com/crawsible/crawsibot/eventinterp"

type FakeInterp struct {
	BeginInterpretingCalls  int
	BeginInterpretingClient eventinterp.Forwarder
}

func (a *FakeInterp) BeginInterpreting(fwdr eventinterp.Forwarder) {
	a.BeginInterpretingCalls += 1
	a.BeginInterpretingClient = fwdr
}
