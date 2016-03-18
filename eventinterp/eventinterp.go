package eventinterp

import "github.com/crawsible/crawsibot/irc/models"

type Enroller interface {
	EnrollForMsgs(cmd string) chan *models.Message
}

type Interp interface {
	BeginInterpreting(enrl Enroller)
	RegisterForInterp(rcvr LoginRcvr)
}

type EventInterp struct {
	LoginInterp Interp
}

func New() *EventInterp {
	return &EventInterp{
		LoginInterp: &LoginInterp{},
	}
}

func (e *EventInterp) BeginInterpreting(enlr Enroller) {
	e.LoginInterp.BeginInterpreting(enlr)
}

func (e *EventInterp) RegisterForLogin(rcvr LoginRcvr) {
	e.LoginInterp.RegisterForInterp(rcvr)
}
