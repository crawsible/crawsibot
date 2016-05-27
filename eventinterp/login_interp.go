package eventinterp

import (
	"github.com/crawsible/crawsibot/eventinterp/event"
	"github.com/crawsible/crawsibot/irc/message"
)

type LoginInterp struct {
	BaseInterp

	MsgCh chan *message.Message
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
