package eventinterp

type LoginAnalyst struct{}

func (l *LoginAnalyst) BeginInterpreting(fwdr Forwarder) {
	fwdr.EnrollForMsgs(l, "RPL_ENDOFMOTD")
}

func (l *LoginAnalyst) RcvMsg(nick, fprms, prms string) {}
