package mocks

import (
	"io"

	"github.com/crawsible/crawsibot/irc"
)

type FakeSender struct {
	StartSendingCalls  int
	StartSendingWriter io.Writer
	ReturnCh           chan *irc.Message

	rcvdMsgs []*irc.Message
}

func (s *FakeSender) StartSending(wtr io.Writer) chan *irc.Message {
	s.StartSendingCalls += 1
	s.StartSendingWriter = wtr

	if s.ReturnCh == nil {
		return make(chan *irc.Message, 1000)
	} else {
		go func() {
			for msg := range s.ReturnCh {
				s.rcvdMsgs = append(s.rcvdMsgs, msg)
			}
		}()

		return s.ReturnCh
	}
}

func (s *FakeSender) ReceivedOverChan() []*irc.Message {
	return s.rcvdMsgs
}
