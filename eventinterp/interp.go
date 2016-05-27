package eventinterp

import "github.com/crawsible/crawsibot/eventinterp/event"

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
			close(eventCh)
		}
	}
}