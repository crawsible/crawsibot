package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeCipher struct {
	EncodeCalls    int
	EncodeMessages []*irc.Message
	EncodeReturns  []string
}

func (c *FakeCipher) Encode(msg *irc.Message) string {
	res := ""
	if c.EncodeCalls < len(c.EncodeReturns) {
		res = c.EncodeReturns[c.EncodeCalls]
	}

	c.EncodeCalls += 1
	c.EncodeMessages = append(c.EncodeMessages, msg)

	return res
}

func (c *FakeCipher) Decode(msgStr string) (*irc.Message, error) {
	return &irc.Message{}, nil
}
