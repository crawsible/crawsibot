package irc

import "io"

type MessageSender struct {
	SendCh chan *Message
}

type Encoder interface {
	Encode(*Message) string
}

func (s *MessageSender) StartSending(wtr io.Writer, ecdr Encoder) {
	s.SendCh = make(chan *Message, 90)
	go func() {
		for msg := range s.SendCh {
			wtr.Write([]byte(ecdr.Encode(msg)))
		}
	}()
}

func (s *MessageSender) Send(cmd, fprms, prms string) {
	s.SendCh <- &Message{"", cmd, fprms, prms}
}
