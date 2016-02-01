package main

import (
	"flag"
	"os"

	"github.com/crawsible/crawsibot/config"
	"github.com/crawsible/crawsibot/irc"
)

var (
	cfg    *config.Config
	client *irc.IRC
)

func init() {
	cfg = &config.Config{}
	client = irc.New()
}

func main() {
	cfg.MakeFlags(flag.NewFlagSet("config", flag.PanicOnError), os.Args[1:])
	client.Connect(cfg)
}
