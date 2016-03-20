package irc

import (
	"io"

	"github.com/crawsible/crawsibot/irc/message"
)

type MessageSender struct {
	SendCh chan *message.Message
}

type Encoder interface {
	Encode(*message.Message) string
}

func (s *MessageSender) StartSending(wtr io.Writer, ecdr Encoder) {
	s.SendCh = make(chan *message.Message, 90)
	go func() {
		for msg := range s.SendCh {
			wtr.Write([]byte(ecdr.Encode(msg)))
		}
	}()
}

func (s *MessageSender) Send(cmd, fprms, prms string) {
	s.SendCh <- &message.Message{"", cmd, fprms, prms}
}
