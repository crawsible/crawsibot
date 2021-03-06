package mocks

import "github.com/crawsible/crawsibot/irc/message"

type FakeCipher struct {
	EncodeMessages []*message.Message
	EncodeStrings  []string
	encodeCalls    int

	DecodeStrings  []string
	DecodeMessages []*message.Message
	decodeCalls    int
}

func (c *FakeCipher) Encode(msg *message.Message) (str string) {
	c.encodeCalls += 1
	c.EncodeMessages = append(c.EncodeMessages, msg)

	if len(c.EncodeStrings) == 0 {
		return
	}

	str = c.EncodeStrings[0]
	c.EncodeStrings = c.EncodeStrings[1:]
	return
}

func (c *FakeCipher) EncodeCalls() int {
	return c.encodeCalls
}

func (c *FakeCipher) Decode(str string) (msg *message.Message, err error) {
	c.decodeCalls += 1
	c.DecodeStrings = append(c.DecodeStrings, str)

	msg = &message.Message{}
	if len(c.DecodeMessages) == 0 {
		return
	}

	msg = c.DecodeMessages[0]
	c.DecodeMessages = c.DecodeMessages[1:]
	return
}

func (c *FakeCipher) DecodeCalls() int {
	return c.decodeCalls
}
