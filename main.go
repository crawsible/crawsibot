package main

import (
	"flag"
	"os"

	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/eventinterp"
	"github.com/crawsible/crawsibot/irc"
)

var (
	cfg        *config.Config
	client     *irc.IRC
	controller *eventinterp.EventInterp
)

func init() {
	cfg = &config.Config{}
	client = irc.New()
	controller = eventinterp.New()
}

func main() {
	cfg.MakeFlags(flag.NewFlagSet("config", flag.PanicOnError), os.Args[1:])
	controller.BeginInterpreting(client)
	client.Connect(cfg)

	keepaliveCh := make(chan struct{})
	<-keepaliveCh
}
