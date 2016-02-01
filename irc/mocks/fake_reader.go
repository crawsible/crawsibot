package mocks

import "errors"

type FakeReader struct {
	ReadStringByte  byte
	RdStrCh         chan string
	readStringCalls int
}

func (c *FakeReader) ReadString(b byte) (str string, err error) {
	str = <-c.RdStrCh
	if str == "EOF" {
		err = errors.New("MOCK READSTRING ERR")
	}

	c.readStringCalls += 1
	c.ReadStringByte = b
	return
}

func (c *FakeReader) ReadStringCalls() int {
	return c.readStringCalls
}
