package main

import (
	"flag"
	"os"

	"github.com/crawsible/crawsibot/config"
)

func main() {
	cfg := &config.Config{}
	cfg.MakeFlags(flag.NewFlagSet("config", flag.PanicOnError), os.Args[1:])
}
