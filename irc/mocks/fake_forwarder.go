package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeForwarder struct {
	StartForwardingCalls  int
	StartForwardingReader irc.ReadStringer
	StartForwardingDcdr   irc.Decoder

	EnrollForPINGCalls     int
	EnrollForPINGRcvr irc.PINGRcvr
}

func (f *FakeForwarder) StartForwarding(rsr irc.ReadStringer, dcdr irc.Decoder) {
	f.StartForwardingCalls += 1
	f.StartForwardingReader = rsr
	f.StartForwardingDcdr = dcdr
}

func (f *FakeForwarder) EnrollForPING(rcp irc.PINGRcvr) {
	f.EnrollForPINGCalls += 1
	f.EnrollForPINGRcvr = rcp
}
