package eventinterp

import "github.com/crawsible/crawsibot/irc"

type Forwarder interface {
	EnrollForMsgs(mrc irc.MsgRcvr, cmd string)
}

type EventInterp struct{}

func New() *EventInterp {
	return &EventInterp{}
}

func (e *EventInterp) BeginInterpreting(fwdr Forwarder) {
	fwdr.EnrollForMsgs(e, "RPL_ENDOFMOTD")
}

func (e *EventInterp) RcvMsg(nick, fprms, prms string) {}
