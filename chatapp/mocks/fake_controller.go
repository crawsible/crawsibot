package mocks

import "github.com/crawsible/crawsibot/eventinterp"

type FakeController struct{}

func (c *FakeController) RegisterForLogin(rcvr eventinterp.LoginRcvr) {}
