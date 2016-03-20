package mocks

import (
	"github.com/crawsible/crawsibot/irc"
	"github.com/crawsible/crawsibot/irc/message"
)

type FakeForwarder struct {
	StartForwardingCalls  int
	StartForwardingReader irc.ReadStringer
	StartForwardingDcdr   irc.Decoder

	EnrollForMsgCalls      int
	EnrollForMsgReturnChan chan *message.Message
	EnrollForMsgCmd        string
}

func (f *FakeForwarder) StartForwarding(rsr irc.ReadStringer, dcdr irc.Decoder) {
	f.StartForwardingCalls += 1
	f.StartForwardingReader = rsr
	f.StartForwardingDcdr = dcdr
}

func (f *FakeForwarder) EnrollForMsgs(cmd string) chan *message.Message {
	f.EnrollForMsgCalls += 1
	f.EnrollForMsgCmd = cmd
	return f.EnrollForMsgReturnChan
}
