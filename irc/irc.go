package irc

import (
	"bufio"
	"io"
	"net"

	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/irc/models"
)

type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}

type Cipher interface {
	Encoder
	Decoder
}

type Sender interface {
	StartSending(wtr io.Writer, ecdr Encoder)
	Send(cmd, fprms, prms string)
}

type Forwarder interface {
	StartForwarding(rsr ReadStringer, dcdr Decoder)
	EnrollForMsgs(cmd string) chan *models.Message
}

type Ponger interface {
	StartPonging(msgr Messenger)
}

type Messenger interface {
	EnrollForMsgs(cmd string) chan *models.Message
	Send(cmd, fprms, prms string)
}

type IRC struct {
	Dialer    Dialer
	Cipher    Cipher
	Sender    Sender
	Forwarder Forwarder
	Ponger    Ponger
}

func New() *IRC {
	return &IRC{
		Dialer:    &net.Dialer{},
		Cipher:    &MessageCipher{},
		Sender:    &MessageSender{},
		Forwarder: &MessageForwarder{make(map[string][]MsgRcvr)},
		Ponger:    &ServerPonger{},
	}
}

func (i *IRC) Connect(cfg *config.Config) {
	conn, err := i.Dialer.Dial("tcp", cfg.Address)
	if err != nil {
		panic(err.Error())
	}

	i.Sender.StartSending(conn, i.Cipher)
	i.Ponger.StartPonging(i)
	i.Forwarder.StartForwarding(bufio.NewReader(conn), i.Cipher)

	i.Sender.Send("PASS", cfg.Password, "")
	i.Sender.Send("NICK", cfg.Nick, "")
	i.Sender.Send("CAP", "REQ", "twitch.tv/membership")
}

func (i *IRC) Send(cmd, fprms, prms string) {
	i.Sender.Send(cmd, fprms, prms)
}

func (i *IRC) EnrollForMsgs(cmd string) chan *models.Message {
	return i.Forwarder.EnrollForMsgs(cmd)
}
