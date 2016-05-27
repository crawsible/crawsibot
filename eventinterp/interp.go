package eventinterp

import "github.com/crawsible/crawsibot/eventinterp/event"

type BaseInterp struct {
	EventChs []chan *event.Event
}

func (in *BaseInterp) RegisterForInterp(eventCh chan *event.Event) {
	in.EventChs = append(in.EventChs, eventCh)
}
