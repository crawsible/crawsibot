package irc

import (
	"fmt"
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
	conn, _ := i.Dialer.Dial("tcp", cfg.Address)
	defer conn.Close()

	fmt.Fprintf(conn, "PASS %s\r\nNICK %s\r\n", cfg.Password, cfg.Nick)
}
