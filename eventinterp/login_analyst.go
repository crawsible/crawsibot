package eventinterp

type LoginRcvr interface {
	LoggedIn()
}

type LoginAnalyst struct {
	EventCh    chan struct{}
	LoginRcvrs []LoginRcvr
}

func (l *LoginAnalyst) BeginInterpreting(fwdr Forwarder) {
	l.EventCh = make(chan struct{}, 1)
	fwdr.EnrollForMsgs(l, "RPL_ENDOFMOTD")

	go l.listenForLogin()
}

func (l *LoginAnalyst) listenForLogin() {
	<-l.EventCh
	for _, rcvr := range l.LoginRcvrs {
		rcvr.LoggedIn()
	}
}

func (l *LoginAnalyst) RcvMsg(n, f, p string) {
	l.EventCh <- struct{}{}
}
