package irc

import (
	"io"
	"net"

	"github.com/crawsible/crawsibot/config"
)

type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}

type IRCSender interface {
	StartSending(wtr io.Writer) chan *Message
}

type IRC struct {
	Dialer Dialer
	Sender IRCSender

	sendCh chan *Message
}

func New() *IRC {
	return &IRC{
		Dialer: &net.Dialer{},
		Sender: NewSender(),
	}
}

func (i *IRC) Connect(cfg *config.Config) {
	conn, _ := i.Dialer.Dial("tcp", cfg.Address)

	i.sendCh = i.Sender.StartSending(conn)
	i.validate(cfg)
}

func (i *IRC) validate(cfg *config.Config) {
	i.sendCh <- &Message{Command: "PASS", FirstParams: cfg.Password}
	i.sendCh <- &Message{Command: "NICK", FirstParams: cfg.Nick}
	i.sendCh <- &Message{
		Command:     "CAP",
		FirstParams: "REQ",
		Params:      "twitch.tv/membership",
	}
}
