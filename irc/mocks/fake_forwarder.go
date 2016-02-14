package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeForwarder struct {
	StartForwardingCalls  int
	StartForwardingReader irc.ReadStringer

	EnrollForPINGCalls     int
	EnrollForPINGRecipient irc.PINGRecipient
}

func (f *FakeForwarder) StartForwarding(rsr irc.ReadStringer) {
	f.StartForwardingCalls += 1
	f.StartForwardingReader = rsr
}

func (f *FakeForwarder) EnrollForPING(rcp irc.PINGRecipient) {
	f.EnrollForPINGCalls += 1
	f.EnrollForPINGRecipient = rcp
}
