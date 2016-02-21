package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeClient struct {
	EnrollForMsgsCalls int
	EnrollForMsgsRcvr  irc.MsgRcvr
	EnrollForMsgsCmd   string
}

func (c *FakeClient) EnrollForMsgs(mrc irc.MsgRcvr, cmd string) {
	c.EnrollForMsgsCalls += 1
	c.EnrollForMsgsRcvr = mrc
	c.EnrollForMsgsCmd = cmd
}
