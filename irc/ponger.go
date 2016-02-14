package irc

type Ponger struct {
	PingCh chan string
}

func (p *Ponger) StartPonging(msgr Messenger) {
	p.PingCh = make(chan string)
	msgr.EnrollForPING(p)

	go func() {
		for server := range p.PingCh {
			msgr.Send("PONG", "", server)
		}
	}()
}

func (p *Ponger) RcvPING(nick, fprms, prms string) {
	p.PingCh <- prms
}
