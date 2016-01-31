package irc

import (
	"net"

	"github.com/crawsible/crawsibot/config"
)

type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}

type IRC struct {
	Dialer Dialer
}

func New() *IRC {
	return &IRC{
		Dialer: &net.Dialer{},
	}
}

func (i *IRC) Connect(cfg *config.Config) {
	i.Dialer.Dial("tcp", cfg.Address)
}
