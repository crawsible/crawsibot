package mocks

import "github.com/crawsible/crawsibot/eventinterp/event"

type FakeRegistrar struct {
	EnrollForEventsCalls      int
	EnrollForEventsTypes      []event.Type
	EnrollForEventsReturnChan chan *event.Event
}

func (r *FakeRegistrar) EnrollForEvents(eventTypes ...event.Type) chan *event.Event {
	r.EnrollForEventsCalls += 1
	r.EnrollForEventsTypes = eventTypes

	return r.EnrollForEventsReturnChan
}
