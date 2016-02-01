package mocks

type FakeRecipient struct {
	RcvPINGNick  string
	RcvPINGFprms string
	RcvPINGPrms  string
	rcvPINGCalls int
}

func (r *FakeRecipient) RcvPING(nick, fprms, prms string) {
	r.rcvPINGCalls += 1

	r.RcvPINGNick = nick
	r.RcvPINGFprms = fprms
	r.RcvPINGPrms = prms
}

func (r *FakeRecipient) RcvPINGCalls() int {
	return r.rcvPINGCalls
}
