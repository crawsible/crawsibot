package irc

import (
	"io"

	"github.com/crawsible/crawsibot/irc/models"
)

type MessageSender struct {
	SendCh chan *models.Message
}

type Encoder interface {
	Encode(*models.Message) string
}

func (s *MessageSender) StartSending(wtr io.Writer, ecdr Encoder) {
	s.SendCh = make(chan *models.Message, 90)
	go func() {
		for msg := range s.SendCh {
			wtr.Write([]byte(ecdr.Encode(msg)))
		}
	}()
}

func (s *MessageSender) Send(cmd, fprms, prms string) {
	s.SendCh <- &models.Message{"", cmd, fprms, prms}
}
