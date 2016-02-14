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
	StartForwarding(ReadStringer)
	EnrollForPING(PINGRecipient)
}

type IRCPonger interface {
	StartPonging(msgr Messenger)
}

type Messenger interface {
	EnrollForPING(PINGRecipient)
	Send(cmd, fprms, prms string)
}

type IRC struct {
	Dialer    Dialer
	Sender    IRCSender
	Forwarder IRCForwarder
	Ponger    IRCPonger
}

func New() *IRC {
	return &IRC{
		Dialer:    &net.Dialer{},
		Sender:    NewSender(),
		Forwarder: NewForwarder(),
		Ponger:    &Ponger{},
	}
}

func (i *IRC) Connect(cfg *config.Config) {
	conn, err := i.Dialer.Dial("tcp", cfg.Address)
	if err != nil {
		panic(err.Error())
	}

	i.Sender.StartSending(conn)
	i.Ponger.StartPonging(i)
	i.Forwarder.StartForwarding(bufio.NewReader(conn))

	i.Sender.Send("PASS", cfg.Password, "")
	i.Sender.Send("NICK", cfg.Nick, "")
	i.Sender.Send("CAP", "REQ", "twitch.tv/membership")
}

func (i *IRC) Send(cmd, fprms, prms string) {
	i.Sender.Send(cmd, fprms, prms)
}

func (i *IRC) EnrollForPING(rcp PINGRecipient) {
	i.Forwarder.EnrollForPING(rcp)
}
