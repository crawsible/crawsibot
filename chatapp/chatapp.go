package chatapp

import "github.com/crawsible/crawsibot/eventinterp"

type Registrar interface {
	RegisterForLogin(rcvr eventinterp.LoginRcvr)
}

type App interface {
	BeginChatting(rgsr Registrar)
}

type ChatApp struct {
	JoinChannelApp App
}

func New() *ChatApp {
	return &ChatApp{
		JoinChannelApp: &JoinChannelApp{},
	}
}

func (c *ChatApp) BeginChatting(rgsr Registrar) {
	c.JoinChannelApp.BeginChatting(rgsr)
}
