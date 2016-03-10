package eventinterp

import "github.com/crawsible/crawsibot/irc"

type Forwarder interface {
	EnrollForMsgs(mrc irc.MsgRcvr, cmd string)
}

type Interp interface {
	BeginInterpreting(fwdr Forwarder)
}

type EventInterp struct {
	LoginInterp Interp
}

func New() *EventInterp {
	return &EventInterp{
		LoginInterp: &LoginInterp{},
	}
}

func (e *EventInterp) BeginInterpreting(fwdr Forwarder) {
	e.LoginInterp.BeginInterpreting(fwdr)
}

func (e *EventInterp) RcvMsg(nick, fprms, prms string) {}
