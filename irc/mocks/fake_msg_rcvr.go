package mocks

type FakeMsgRcvr struct {
	RcvMsgNick  string
	RcvMsgFprms string
	RcvMsgPrms  string
	rcvMsgCalls int
}

func (r *FakeMsgRcvr) RcvMsg(nick, fprms, prms string) {
	r.rcvMsgCalls += 1

	r.RcvMsgNick = nick
	r.RcvMsgFprms = fprms
	r.RcvMsgPrms = prms
}

func (r *FakeMsgRcvr) RcvMsgCalls() int {
	return r.rcvMsgCalls
}
