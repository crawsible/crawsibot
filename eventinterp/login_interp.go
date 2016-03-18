package eventinterp

import "github.com/crawsible/crawsibot/irc/models"

type LoginRcvr interface {
	LoggedIn()
}

type LoginInterp struct {
	EventCh    chan *models.Message
	LoginRcvrs []LoginRcvr
}

func (l *LoginInterp) RegisterForInterp(rcvr LoginRcvr) {
	l.LoginRcvrs = append(l.LoginRcvrs, rcvr)
}

func (l *LoginInterp) BeginInterpreting(enlr Enroller) {
	l.EventCh = enlr.EnrollForMsgs("RPL_ENDOFMOTD")
	go l.listenForLogin()
}

func (l *LoginInterp) listenForLogin() {
	<-l.EventCh
	for _, rcvr := range l.LoginRcvrs {
		rcvr.LoggedIn()
	}
}
