package mocks

import "github.com/crawsible/crawsibot/irc"

type FakePonger struct {
	StartPongingCalls int
	StartPongingMsgr  irc.Messenger

	FakeForwarder       *FakeForwarder
	ForwarderHadStarted bool
}

func (f *FakePonger) StartPonging(msgr irc.Messenger) {
	f.StartPongingCalls += 1
	f.StartPongingMsgr = msgr

	f.ForwarderHadStarted = (f.FakeForwarder.StartForwardingCalls > 0)
}
