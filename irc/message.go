package irc

type Message struct {
	NickOrSrvname string
	Command       string
	FirstParams   string
	Params        string
}
