package mocks

type FakePonger struct{}

func (f *FakePonger) RcvPING(nick, fprms, prms string) {}
