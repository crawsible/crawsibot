package irc

type Message struct {
	NickOrSrvname string
	Command       string
	FirstParam    string
	Params        string
}
