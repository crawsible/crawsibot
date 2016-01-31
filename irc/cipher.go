package irc

import (
	"errors"
	"fmt"
	"regexp"
)

// see https://tools.ietf.org/html/rfc1459#section-2.3
// Exceptions have been afforded for Twitch compatibility

type MessageCipher struct{}

var msgRE *regexp.Regexp = regexp.MustCompile(
	`\A` +
		`(?:\:(?P<srvname>[a-zA-Z][^!\s]*)\S*\s+)?` +
		`(?P<cmd>[a-zA-Z]+|[0-9]{3})\s+` +
		`(?P<firstprms>[^:]+)` +
		`(?:\s+\:(?P<params>.*))?` +
		`\r\n\z`,
)

func (c *MessageCipher) Decode(msgStr string) (*Message, error) {
	match := msgRE.FindStringSubmatch(msgStr)
	if len(match) == 0 {
		errMsg := fmt.Sprintf(
			"The message received:\n%s\nis invalid.",
			msgStr,
		)
		return &Message{}, errors.New(errMsg)
	}

	named := getNamedMatch(match)

	return &Message{
		NickOrSrvname: named["srvname"],
		Command:       getStringFor(named["cmd"]),
		FirstParams:   named["firstprms"],
		Params:        named["params"],
	}, nil
}

func (c *MessageCipher) Encode(msg *Message) string {
	params := ""
	if msg.Params != "" {
		params = fmt.Sprintf(" :%s", msg.Params)
	}

	return fmt.Sprintf("%s %s%s\r\n", msg.Command, msg.FirstParams, params)
}

func getNamedMatch(match []string) map[string]string {
	named := make(map[string]string)

	for i, name := range msgRE.SubexpNames() {
		if i != 0 {
			named[name] = match[i]
		}
	}

	return named
}

var codeTable map[string]string = map[string]string{
	"001": "RPL_WELCOME",
	"002": "RPL_YOURHOST",
	"003": "RPL_CREATED",
	"004": "RPL_MYINFO",
	"353": "RPL_NAMREPLY",
	"366": "RPL_ENDOFNAMES",
	"372": "RPL_MOTD",
	"375": "RPL_MOTDSTART",
	"376": "RPL_ENDOFMOTD",
	"421": "ERR_UNKNOWNCOMMAND",
}

func getStringFor(cmd string) string {
	textCmd := codeTable[cmd]
	if textCmd == "" {
		return cmd
	}

	return textCmd
}
