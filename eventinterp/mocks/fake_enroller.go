package mocks

import "github.com/crawsible/crawsibot/irc/models"

type FakeEnroller struct {
	EnrollForMsgsCalls      int
	EnrollForMsgsCmd        string
	EnrollForMsgsReturnChan chan *models.Message
}

func (e *FakeEnroller) EnrollForMsgs(cmd string) chan *models.Message {
	e.EnrollForMsgsCalls += 1
	e.EnrollForMsgsCmd = cmd
	return e.EnrollForMsgsReturnChan
}
