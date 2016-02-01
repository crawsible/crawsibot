package config

type Config struct {
	Address  string
	Nick     string
	Password string
	Channel  string
}

type FlagSet interface {
	StringVar(p *string, name, value, usage string)
	Parse(args []string) error
}

func (c *Config) MakeFlags(fs FlagSet, args []string) {
	fs.StringVar(&(c.Address), "a", "irc.twitch.tv:6667", "The IRC server address (<hostname>:<port>)")
	fs.StringVar(&(c.Nick), "n", "", "Your IRC nickname")
	fs.StringVar(&(c.Password), "p", "", "Your IRC password")
	fs.StringVar(&(c.Channel), "c", "", "The channel you're connecting to (omit '#')")

	if err := fs.Parse(args); err != nil {
		panic(err.Error())
	}
}
