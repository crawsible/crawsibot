package mocks

import (
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/eventinterp/event"
)

type FakeInterp struct {
	BeginInterpretingCalls    int
	BeginInterpretingEnroller eventinterp.Enroller

	RegisterForInterpCalls int
	RegisterForInterpChan  chan *event.Event

	UnsubscribeCalls int
	UnsubscribeChan  chan *event.Event
}

func (a *FakeInterp) BeginInterpreting(enlr eventinterp.Enroller) {
	a.BeginInterpretingCalls += 1
	a.BeginInterpretingEnroller = enlr
}

func (a *FakeInterp) RegisterForInterp(ch chan *event.Event) {
	a.RegisterForInterpCalls += 1
	a.RegisterForInterpChan = ch
}

func (a *FakeInterp) Unsubscribe(ch chan *event.Event) {
	a.UnsubscribeCalls += 1
	a.UnsubscribeChan = ch
}
