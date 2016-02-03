package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeForwarder struct {
	StartForwardingCalls  int
	StartForwardingReader irc.ReadStringer

	EnrollForPINGCalls          int
	EnrollForPINGPonger         irc.PINGRecipient
	EnrollForPINGPreviousStarts int
}

func (f *FakeForwarder) StartForwarding(rsr irc.ReadStringer) {
	f.StartForwardingCalls += 1
	f.StartForwardingReader = rsr
}

func (f *FakeForwarder) EnrollForPING(rcp irc.PINGRecipient) {
	f.EnrollForPINGCalls += 1
	f.EnrollForPINGPonger = rcp
	f.EnrollForPINGPreviousStarts = f.StartForwardingCalls
}
