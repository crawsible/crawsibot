package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeMessenger struct {
	EnrollForMsgsCalls int
	EnrollForMsgsRcvr  irc.MsgRcvr
	EnrollForMsgsCmd   string

	SendArgs  [][]string
	sendCalls int
}

func (m *FakeMessenger) EnrollForMsgs(mrc irc.MsgRcvr, cmd string) {
	m.EnrollForMsgsCalls += 1
	m.EnrollForMsgsRcvr = mrc
	m.EnrollForMsgsCmd = cmd
}

func (m *FakeMessenger) Send(cmd, fprms, prms string) {
	m.sendCalls += 1
	m.SendArgs = append(m.SendArgs, []string{cmd, fprms, prms})
}

func (m *FakeMessenger) SendCalls() int {
	return m.sendCalls
}
