package irc

import "github.com/crawsible/crawsibot/irc/models"

type ServerPonger struct {
	PingCh chan *models.Message
}

func (p *ServerPonger) StartPonging(msgr Messenger) {
	p.PingCh = msgr.EnrollForMsgs("PING")

	go func() {
		for msg := range p.PingCh {
			msgr.Send("PONG", "", msg.Params)
		}
	}()
}
