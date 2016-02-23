package mocks

import "github.com/crawsible/crawsibot/eventinterp"

type FakeAnalyst struct {
	BeginInterpretingCalls  int
	BeginInterpretingClient eventinterp.Forwarder
}

func (a *FakeAnalyst) BeginInterpreting(fwdr eventinterp.Forwarder) {
	a.BeginInterpretingCalls += 1
	a.BeginInterpretingClient = fwdr
}
