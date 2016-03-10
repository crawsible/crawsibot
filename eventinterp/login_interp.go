package eventinterp

type LoginRcvr interface {
	LoggedIn()
}

type LoginInterp struct {
	EventCh    chan struct{}
	LoginRcvrs []LoginRcvr
}

func (l *LoginInterp) BeginInterpreting(fwdr Forwarder) {
	l.EventCh = make(chan struct{}, 1)
	fwdr.EnrollForMsgs(l, "RPL_ENDOFMOTD")

	go l.listenForLogin()
}

func (l *LoginInterp) listenForLogin() {
	<-l.EventCh
	for _, rcvr := range l.LoginRcvrs {
		rcvr.LoggedIn()
	}
}

func (l *LoginInterp) RcvMsg(n, f, p string) {
	l.EventCh <- struct{}{}
}
