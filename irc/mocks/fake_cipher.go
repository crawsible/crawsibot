package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeCipher struct {
	encodeCalls    int
	EncodeMessages []*irc.Message
	EncodeReturns  []string
}

func (c *FakeCipher) Encode(msg *irc.Message) string {
	res := ""
	if c.encodeCalls < len(c.EncodeReturns) {
		res = c.EncodeReturns[c.encodeCalls]
	}

	c.encodeCalls += 1
	c.EncodeMessages = append(c.EncodeMessages, msg)

	return res
}

func (c *FakeCipher) EncodeCalls() int {
	return c.encodeCalls
}

func (c *FakeCipher) Decode(msgStr string) (*irc.Message, error) {
	return &irc.Message{}, nil
}
