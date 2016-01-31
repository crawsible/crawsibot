package irc

import "regexp"

type Cipher struct{}

var msgRE *regexp.Regexp = regexp.MustCompile(
	`\A` +
		`(?:\:(?P<srvname>[a-zA-Z][^!\s]*)\S*\s+)?` +
		`(?P<cmd>[a-zA-Z]+|[0-9]{3})\s+` +
		`(?P<firstprm>[^:]+)` +
		`(?:\s+\:(?P<params>.*))?` +
		`\r\n\z`,
)

func (c *Cipher) Decode(msgStr string) *Message {
	match := msgRE.FindStringSubmatch(msgStr)
	named := getNamedMatch(match)

	return &Message{
		NickOrSrvname: named["srvname"],
		Command:       named["cmd"],
		FirstParam:    named["firstprm"],
		Params:        named["params"],
	}
}

func getNamedMatch(match []string) map[string]string {
	named := map[string]string{}

	for i, name := range msgRE.SubexpNames() {
		if i != 0 {
			named[name] = match[i]
		}
	}

	return named
}
