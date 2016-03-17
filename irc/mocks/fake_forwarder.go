package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeForwarder struct {
	StartForwardingCalls  int
	StartForwardingReader irc.ReadStringer
	StartForwardingDcdr   irc.Decoder

	EnrollForMsgCalls int
	EnrollForMsgChan  chan map[string]string
	EnrollForMsgCmd   string
}

func (f *FakeForwarder) StartForwarding(rsr irc.ReadStringer, dcdr irc.Decoder) {
	f.StartForwardingCalls += 1
	f.StartForwardingReader = rsr
	f.StartForwardingDcdr = dcdr
}

func (f *FakeForwarder) EnrollForMsgs(rcvCh chan map[string]string, cmd string) {
	f.EnrollForMsgCalls += 1
	f.EnrollForMsgChan = rcvCh
	f.EnrollForMsgCmd = cmd
}
