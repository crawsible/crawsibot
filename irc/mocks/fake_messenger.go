package mocks

import "github.com/crawsible/crawsibot/irc/message"

type FakeMessenger struct {
	EnrollForMsgsCalls      int
	EnrollForMsgsReturnChan chan *message.Message
	EnrollForMsgsCmd        string

	SendArgs  [][]string
	sendCalls int
}

func (m *FakeMessenger) EnrollForMsgs(cmd string) chan *message.Message {
	m.EnrollForMsgsCalls += 1
	m.EnrollForMsgsCmd = cmd
	return m.EnrollForMsgsReturnChan
}

func (m *FakeMessenger) Send(cmd, fprms, prms string) {
	m.sendCalls += 1
	m.SendArgs = append(m.SendArgs, []string{cmd, fprms, prms})
}

func (m *FakeMessenger) SendCalls() int {
	return m.sendCalls
}
