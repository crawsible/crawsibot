package eventinterp

import (
	"regexp"

	"github.com/crawsible/crawsibot/eventinterp/event"
	"github.com/crawsible/crawsibot/irc/message"
)

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
