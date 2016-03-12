package mocks

import (
	"github.com/crawsible/crawsibot/chatapp"
	"github.com/crawsible/crawsibot/config"
)

type FakeApp struct {
	BeginChattingCalls     int
	BeginChattingRegistrar chatapp.Registrar
	BeginChattingSender    chatapp.Sender
	BeginChattingCfg       *config.Config
}

func (a *FakeApp) BeginChatting(rgsr chatapp.Registrar, sndr chatapp.Sender, cfg *config.Config) {
	a.BeginChattingCalls += 1
	a.BeginChattingRegistrar = rgsr
	a.BeginChattingSender = sndr
	a.BeginChattingCfg = cfg
}
