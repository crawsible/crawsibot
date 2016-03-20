package irc

import "github.com/crawsible/crawsibot/irc/message"

type ServerPonger struct {
	PingCh chan *message.Message
}

func (p *ServerPonger) StartPonging(msgr Messenger) {
	p.PingCh = msgr.EnrollForMsgs("PING")

	go func() {
		for msg := range p.PingCh {
			msgr.Send("PONG", "", msg.Params)
		}
	}()
}
