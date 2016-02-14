package irc

type ServerPonger struct {
	PingCh chan string
}

func (p *ServerPonger) StartPonging(msgr Messenger) {
	p.PingCh = make(chan string)
	msgr.EnrollForPING(p)

	go func() {
		for server := range p.PingCh {
			msgr.Send("PONG", "", server)
		}
	}()
}

func (p *ServerPonger) RcvPING(nick, fprms, prms string) {
	p.PingCh <- prms
}
