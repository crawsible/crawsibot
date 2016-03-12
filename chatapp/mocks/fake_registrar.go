package mocks

import "github.com/crawsible/crawsibot/eventinterp"

type FakeRegistrar struct{}

func (c *FakeRegistrar) RegisterForLogin(rcvr eventinterp.LoginRcvr) {}
