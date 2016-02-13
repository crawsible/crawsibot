package mocks

import (
	"io"

	"github.com/crawsible/crawsibot/irc"
)

type FakeSender struct {
	StartSendingCalls  int
	StartSendingWriter io.Writer
	SendCh             chan *irc.Message

	SendCalls int
	SendArgs  [][]string
}

func (s *FakeSender) StartSending(wtr io.Writer) {
	s.StartSendingCalls += 1
	s.StartSendingWriter = wtr
}

func (s *FakeSender) Send(cmd, fprms, prms string) {
	s.SendCalls += 1
	s.SendArgs = append(s.SendArgs, []string{cmd, fprms, prms})
}
