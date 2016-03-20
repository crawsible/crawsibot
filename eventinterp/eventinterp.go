package eventinterp

import (
	"github.com/crawsible/crawsibot/eventinterp/event"
	"github.com/crawsible/crawsibot/irc/models"
)

type Enroller interface {
	EnrollForMsgs(cmd string) chan *models.Message
}

type Interp interface {
	BeginInterpreting(enrl Enroller)
	RegisterForInterp(eventCh chan *event.Event)
}

type EventInterp struct {
	Interps map[event.Type]Interp
}

func New() *EventInterp {
	return &EventInterp{
		Interps: map[event.Type]Interp{
			event.Login: &LoginInterp{},
		},
	}
}

func (e *EventInterp) BeginInterpreting(enrl Enroller) {
	for _, interp := range e.Interps {
		interp.BeginInterpreting(enrl)
	}
}

func (e *EventInterp) EnrollForEvents(eventTypes ...event.Type) chan *event.Event {
	eventCh := make(chan *event.Event, 1)
	for _, eventType := range eventTypes {
		e.Interps[eventType].RegisterForInterp(eventCh)
	}

	return eventCh
}
