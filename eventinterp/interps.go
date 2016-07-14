package eventinterp

import (
	"regexp"

	"github.com/crawsible/crawsibot/eventinterp/event"
	"github.com/crawsible/crawsibot/irc/message"
)

type BaseInterp struct {
	EventChs []chan *event.Event
}

func (in *BaseInterp) RegisterForInterp(eventCh chan *event.Event) {
	in.EventChs = append(in.EventChs, eventCh)
}

func (in *BaseInterp) Unsubscribe(eventCh chan *event.Event) {
	for i, ch := range in.EventChs {
		if ch == eventCh {
			in.EventChs = append(in.EventChs[:i], in.EventChs[i+1:]...)
		}
	}
}

type LoginInterp struct {
	BaseInterp

	MsgCh chan *message.Message
}

func (l *LoginInterp) BeginInterpreting(enlr Enroller) {
	l.MsgCh = enlr.EnrollForMsgs("RPL_ENDOFMOTD")
	go l.listenForLogin()
}

func (l *LoginInterp) listenForLogin() {
	loginEvt := &event.Event{Type: event.Login}

	<-l.MsgCh

	for _, eventCh := range l.EventChs {
		eventCh <- loginEvt
	}
}

type ChannelJoinInterp struct {
	BaseInterp

	MsgCh chan *message.Message
}

func (c *ChannelJoinInterp) BeginInterpreting(enlr Enroller) {
	c.MsgCh = enlr.EnrollForMsgs("RPL_ENDOFNAMES")
	go c.listenForChannelJoin()
}

func (c *ChannelJoinInterp) listenForChannelJoin() {
	joinEvt := &event.Event{
		Type: event.ChannelJoin,
		Data: map[string]string{
			"joinedChannel": channelFromFprms((<-c.MsgCh).FirstParams),
		},
	}

	for _, ch := range c.EventChs {
		ch <- joinEvt
	}
}

var channelJoinFprmsRE *regexp.Regexp = regexp.MustCompile(
	`\A.*#(\w+).*\z`,
)

func channelFromFprms(fprms string) string {
	match := channelJoinFprmsRE.FindStringSubmatch(fprms)
	if len(match) < 2 {
		return ""
	}

	return match[1]
}

type CommandInterp struct {
	BaseInterp
}

func (in CommandInterp) BeginInterpreting(enlr Enroller) {
}
