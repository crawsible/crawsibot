package mocks

import "github.com/crawsible/crawsibot/eventinterp"

type FakeInterp struct {
	BeginInterpretingCalls    int
	BeginInterpretingEnroller eventinterp.Enroller

	RegisterForInterpCalls int
	RegisterForInterpRcvr  eventinterp.LoginRcvr
}

func (a *FakeInterp) BeginInterpreting(enlr eventinterp.Enroller) {
	a.BeginInterpretingCalls += 1
	a.BeginInterpretingEnroller = enlr
}

func (a *FakeInterp) RegisterForInterp(rcvr eventinterp.LoginRcvr) {
	a.RegisterForInterpCalls += 1
	a.RegisterForInterpRcvr = rcvr
}
