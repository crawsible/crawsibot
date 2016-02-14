package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeMessenger struct {
	EnrollForPINGCalls   int
	EnrollForPINGMsgRcvr irc.MsgRcvr

	SendArgs  [][]string
	sendCalls int
}

func (m *FakeMessenger) EnrollForPING(mrc irc.MsgRcvr) {
	m.EnrollForPINGCalls += 1
	m.EnrollForPINGMsgRcvr = mrc
}

func (m *FakeMessenger) Send(cmd, fprms, prms string) {
	m.sendCalls += 1
	m.SendArgs = append(m.SendArgs, []string{cmd, fprms, prms})
}

func (m *FakeMessenger) SendCalls() int {
	return m.sendCalls
}
