package chatapp

import "github.com/crawsible/crawsibot/config"

type JoinChannelApp struct {
	EventCh chan struct{}
}

func (a *JoinChannelApp) BeginChatting(rgsr Registrar, sndr Sender, cfg *config.Config) {
	a.EventCh = make(chan struct{}, 1)
	rgsr.RegisterForLogin(a)

	go a.joinOnLogin(sndr, cfg)
}

func (a *JoinChannelApp) joinOnLogin(sndr Sender, cfg *config.Config) {
	<-a.EventCh
	sndr.Send("JOIN", "#"+cfg.Channel, "")
}

func (a *JoinChannelApp) LoggedIn() {
	a.EventCh <- struct{}{}
}
