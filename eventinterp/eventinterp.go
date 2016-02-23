package eventinterp

import "github.com/crawsible/crawsibot/irc"

type Forwarder interface {
	EnrollForMsgs(mrc irc.MsgRcvr, cmd string)
}

type Analyst interface {
	BeginInterpreting(fwdr Forwarder)
}

type EventInterp struct {
	LoginAnalyst Analyst
}

func New() *EventInterp {
	return &EventInterp{
		LoginAnalyst: &LoginAnalyst{},
	}
}

func (e *EventInterp) BeginInterpreting(fwdr Forwarder) {
	e.LoginAnalyst.BeginInterpreting(fwdr)
}

func (e *EventInterp) RcvMsg(nick, fprms, prms string) {}
