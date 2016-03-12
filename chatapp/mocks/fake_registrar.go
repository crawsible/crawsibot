package mocks

import "github.com/crawsible/crawsibot/eventinterp"

type FakeRegistrar struct {
	RegisterForLoginCalls int
	RegisterForLoginRcvr  eventinterp.LoginRcvr
}

func (r *FakeRegistrar) RegisterForLogin(rcvr eventinterp.LoginRcvr) {
	r.RegisterForLoginCalls += 1
	r.RegisterForLoginRcvr = rcvr
}
