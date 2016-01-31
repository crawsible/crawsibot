package irc

import (
	"fmt"
	"net"

	"github.com/crawsible/crawsibot/config"
)

type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}

type Cipher interface {
	Decode(string) (*Message, error)
	Encode(*Message) string
}

type IRC struct {
	Dialer Dialer
	Cipher Cipher
}

func New() *IRC {
	return &IRC{
		Dialer: &net.Dialer{},
		Cipher: &MessageCipher{},
	}
}

func (i *IRC) Connect(cfg *config.Config) {
	conn, _ := i.Dialer.Dial("tcp", cfg.Address)
	defer conn.Close()

	i.validateWithConn(conn, cfg)
}

func (i *IRC) validateWithConn(conn net.Conn, cfg *config.Config) {
	passMsg := &Message{Command: "PASS", FirstParams: cfg.Password}
	nickMsg := &Message{Command: "NICK", FirstParams: cfg.Nick}
	fmt.Fprintf(conn, "%s%s", i.Cipher.Encode(passMsg), i.Cipher.Encode(nickMsg))
}
