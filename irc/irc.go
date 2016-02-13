package irc

import (
	"bufio"
	"io"
	"net"

	"github.com/crawsible/crawsibot/config"
)

type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}

type IRCSender interface {
	StartSending(wtr io.Writer)
	Send(cmd, fprms, prms string)
}

type IRCForwarder interface {
	EnrollForPING(PINGRecipient)
	StartForwarding(ReadStringer)
}

type IRCPonger interface {
	RcvPING(nick, fprms, prms string)
}

type IRC struct {
	Dialer    Dialer
	Sender    IRCSender
	Forwarder IRCForwarder
	Ponger    PINGRecipient

	sendCh chan *Message
}

func New() *IRC {
	return &IRC{
		Dialer:    &net.Dialer{},
		Sender:    NewSender(),
		Forwarder: &Forwarder{},
		//Ponger:    &Ponger{},
	}
}

func (i *IRC) Connect(cfg *config.Config) {
	conn, _ := i.Dialer.Dial("tcp", cfg.Address)

	i.Sender.StartSending(conn)
	i.Forwarder.EnrollForPING(i.Ponger)
	i.Forwarder.StartForwarding(bufio.NewReader(conn))

	i.Sender.Send("PASS", cfg.Password, "")
	i.Sender.Send("NICK", cfg.Nick, "")
	i.Sender.Send("CAP", "REQ", "twitch.tv/membership")
}

func (i *IRC) Send(cmd, fprms, prms string) {
	i.Sender.Send(cmd, fprms, prms)
}
