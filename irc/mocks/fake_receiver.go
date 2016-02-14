package mocks

type FakeReceiver struct {
	RcvMsgNick  string
	RcvMsgFprms string
	RcvMsgPrms  string
	rcvMsgCalls int
}

func (r *FakeReceiver) RcvMsg(nick, fprms, prms string) {
	r.rcvMsgCalls += 1

	r.RcvMsgNick = nick
	r.RcvMsgFprms = fprms
	r.RcvMsgPrms = prms
}

func (r *FakeReceiver) RcvMsgCalls() int {
	return r.rcvMsgCalls
}
