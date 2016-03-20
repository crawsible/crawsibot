package eventinterp

import (
	"github.com/crawsible/crawsibot/eventinterp/event"
	"github.com/crawsible/crawsibot/irc/models"
)

type LoginInterp struct {
	MsgCh    chan *models.Message
	EventChs []chan *event.Event
}

func (l *LoginInterp) RegisterForInterp(eventCh chan *event.Event) {
	l.EventChs = append(l.EventChs, eventCh)
}

func (l *LoginInterp) BeginInterpreting(enlr Enroller) {
	l.MsgCh = enlr.EnrollForMsgs("RPL_ENDOFMOTD")
	go l.listenForLogin()
}

func (l *LoginInterp) listenForLogin() {
	<-l.MsgCh
	for _, eventCh := range l.EventChs {
		eventCh <- &event.Event{Type: event.Login}
	}
}
