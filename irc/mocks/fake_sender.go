package mocks

import (
	"io"

	"github.com/crawsible/crawsibot/irc"
)

type FakeSender struct {
	StartSendingCalls  int
	StartSendingWriter io.Writer
	StartSendingEcdr   irc.Encoder
	SendCh             chan *irc.Message

	SendCalls int
	SendArgs  [][]string
}

func (s *FakeSender) StartSending(wtr io.Writer, ecdr irc.Encoder) {
	s.StartSendingCalls += 1
	s.StartSendingWriter = wtr
	s.StartSendingEcdr = ecdr
}

func (s *FakeSender) Send(cmd, fprms, prms string) {
	s.SendCalls += 1
	s.SendArgs = append(s.SendArgs, []string{cmd, fprms, prms})
}
