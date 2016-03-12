package chatapp

import (
	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/eventinterp"
)

type Registrar interface {
	RegisterForLogin(rcvr eventinterp.LoginRcvr)
}

type Sender interface {
	Send(cmd, fprms, prms string)
}

type App interface {
	BeginChatting(rgsr Registrar, sndr Sender, cfg *config.Config)
}

type ChatApp struct {
	JoinChannelApp App
}

func New() *ChatApp {
	return &ChatApp{
		JoinChannelApp: &JoinChannelApp{},
	}
}

func (c *ChatApp) BeginChatting(rgsr Registrar, sndr Sender, cfg *config.Config) {
	c.JoinChannelApp.BeginChatting(rgsr, sndr, cfg)
}
