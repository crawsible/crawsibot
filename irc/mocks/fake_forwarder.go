package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeForwarder struct {
	StartForwardingCalls  int
	StartForwardingReader irc.ReadStringer
}

func (f *FakeForwarder) StartForwarding(rsr irc.ReadStringer) {
	f.StartForwardingCalls += 1
	f.StartForwardingReader = rsr
}
