package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeEnroller struct {
	EnrollForMsgsCalls int
	EnrollForMsgsRcvr  irc.MsgRcvr
	EnrollForMsgsCmd   string
}

func (e *FakeEnroller) EnrollForMsgs(mrc irc.MsgRcvr, cmd string) {
	e.EnrollForMsgsCalls += 1
	e.EnrollForMsgsRcvr = mrc
	e.EnrollForMsgsCmd = cmd
}
