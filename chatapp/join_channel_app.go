package chatapp

import "github.com/crawsible/crawsibot/config"

type JoinChannelApp struct{}

func (j *JoinChannelApp) BeginChatting(rgsr Registrar, sndr Sender, cfg *config.Config) {}
