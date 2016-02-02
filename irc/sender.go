package irc

import "io"

type Encoder interface {
	Encode(*Message) string
}

type Sender struct {
	Encoder Encoder
}

func NewSender() *Sender {
	return &Sender{
		Encoder: &MessageCipher{},
	}
}

func (s *Sender) StartSending(wtr io.Writer) chan *Message {
	sendCh := make(chan *Message, 90)
	go func() {
		for msg := range sendCh {
			wtr.Write([]byte(s.Encoder.Encode(msg)))
		}
	}()

	return sendCh
}
