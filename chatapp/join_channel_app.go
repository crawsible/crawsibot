package chatapp

import (
	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/eventinterp/event"
)

type JoinChannelApp struct {
	EventCh  chan *event.Event
	loggedIn bool
}

func (a *JoinChannelApp) BeginChatting(rgsr Registrar, sndr Sender, cfg *config.Config) {
	a.EventCh = rgsr.EnrollForEvents(event.Login, event.ChannelJoin)

	go a.handleEvents(rgsr, sndr, cfg)
}

func (a *JoinChannelApp) handleEvents(rgsr Registrar, sndr Sender, cfg *config.Config) {
	for evt := range a.EventCh {
		if a.hasLoggedIn(evt) {
			a.loggedIn = true
			sndr.Send("JOIN", "#"+cfg.Channel, "")
		} else if a.hasJoinedChannel(evt, cfg.Channel) {
			sndr.Send(
				"PRIVMSG",
				"#"+cfg.Channel,
				"COME WITH ME IF YOU WANT TO LIVE.",
			)
			rgsr.Unsubscribe(a.EventCh)
		}
	}

}

func (a *JoinChannelApp) hasLoggedIn(evt *event.Event) bool {
	return !a.loggedIn && evt.Type == event.Login
}

func (a *JoinChannelApp) hasJoinedChannel(evt *event.Event, channel string) bool {
	return a.loggedIn &&
		evt.Type == event.ChannelJoin &&
		evt.Data["joinedChannel"] == channel
}
