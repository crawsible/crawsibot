package irc

import "io"

type Encoder interface {
	Encode(*Message) string
}

type Sender struct {
	Encoder Encoder
	SendCh  chan *Message
}

func NewSender() *Sender {
	return &Sender{
		Encoder: &MessageCipher{},
	}
}

func (s *Sender) StartSending(wtr io.Writer) {
	s.SendCh = make(chan *Message, 90)
	go func() {
		for msg := range s.SendCh {
			wtr.Write([]byte(s.Encoder.Encode(msg)))
		}
	}()
}

func (s *Sender) Send(cmd, fprms, prms string) {
	s.SendCh <- &Message{"", cmd, fprms, prms}
}
