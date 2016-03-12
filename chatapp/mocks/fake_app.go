package mocks

import "github.com/crawsible/crawsibot/chatapp"

type FakeApp struct {
	BeginChattingCalls     int
	BeginChattingRegistrar chatapp.Registrar
}

func (a *FakeApp) BeginChatting(rgsr chatapp.Registrar) {
	a.BeginChattingCalls += 1
	a.BeginChattingRegistrar = rgsr
}
