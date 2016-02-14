package irc

import "io"

type Encoder interface {
	Encode(*Message) string
}

type Sender struct {
	Encoder Encoder
	SendCh  chan *Message
}

func (s *Sender) StartSending(wtr io.Writer, ecdr Encoder) {
	s.SendCh = make(chan *Message, 90)
	go func() {
		for msg := range s.SendCh {
			wtr.Write([]byte(ecdr.Encode(msg)))
		}
	}()
}

func (s *Sender) Send(cmd, fprms, prms string) {
	s.SendCh <- &Message{"", cmd, fprms, prms}
}
