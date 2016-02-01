package irc

import (
	"fmt"
	"net"
)

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

func (s *Sender) StartSending(conn net.Conn) chan *Message {
	sendCh := make(chan *Message, 90)
	go func() {
		for msg := range sendCh {
			fmt.Fprintf(conn, s.Encoder.Encode(msg))
		}
	}()

	return sendCh
}
