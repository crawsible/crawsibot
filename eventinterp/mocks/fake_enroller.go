package mocks

import "github.com/crawsible/crawsibot/irc/message"

type FakeEnroller struct {
	EnrollForMsgsCalls      int
	EnrollForMsgsCmd        string
	EnrollForMsgsReturnChan chan *message.Message
}

func (e *FakeEnroller) EnrollForMsgs(cmd string) chan *message.Message {
	e.EnrollForMsgsCalls += 1
	e.EnrollForMsgsCmd = cmd
	return e.EnrollForMsgsReturnChan
}
