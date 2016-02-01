package mocks

import (
	"net"

	"github.com/crawsible/crawsibot/irc"
)

type FakeSender struct {
	StartSendingCalls int
	StartSendingConn  net.Conn
	ReturnCh          chan *irc.Message

	rcvdMsgs []*irc.Message
}

func (s *FakeSender) StartSending(conn net.Conn) chan *irc.Message {
	s.StartSendingCalls += 1
	s.StartSendingConn = conn

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
