package chatapp

import (
	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/eventinterp/event"
)

type JoinChannelApp struct {
	EventCh chan *event.Event
}

func (a *JoinChannelApp) BeginChatting(rgsr Registrar, sndr Sender, cfg *config.Config) {
	a.EventCh = rgsr.EnrollForEvents(event.Login)

	go a.joinOnLogin(sndr, cfg)
}

func (a *JoinChannelApp) joinOnLogin(sndr Sender, cfg *config.Config) {
	<-a.EventCh
	sndr.Send("JOIN", "#"+cfg.Channel, "")
}
