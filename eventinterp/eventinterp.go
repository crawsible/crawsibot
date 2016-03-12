package eventinterp

import "github.com/crawsible/crawsibot/irc"

type Enroller interface {
	EnrollForMsgs(mrc irc.MsgRcvr, cmd string)
}

type Interp interface {
	BeginInterpreting(fwdr Enroller)
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
