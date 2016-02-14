package mocks

type FakeReceiver struct {
	RcvPINGNick  string
	RcvPINGFprms string
	RcvPINGPrms  string
	rcvPINGCalls int
}

func (r *FakeReceiver) RcvPING(nick, fprms, prms string) {
	r.rcvPINGCalls += 1

	r.RcvPINGNick = nick
	r.RcvPINGFprms = fprms
	r.RcvPINGPrms = prms
}

func (r *FakeReceiver) RcvPINGCalls() int {
	return r.rcvPINGCalls
}
