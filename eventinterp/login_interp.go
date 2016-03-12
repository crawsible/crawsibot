package eventinterp

type LoginRcvr interface {
	LoggedIn()
}

type LoginInterp struct {
	EventCh    chan struct{}
	LoginRcvrs []LoginRcvr
}

func (l *LoginInterp) RegisterForInterp(rcvr LoginRcvr) {
	l.LoginRcvrs = append(l.LoginRcvrs, rcvr)
}

func (l *LoginInterp) BeginInterpreting(enlr Enroller) {
	l.EventCh = make(chan struct{}, 1)
	enlr.EnrollForMsgs(l, "RPL_ENDOFMOTD")

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
