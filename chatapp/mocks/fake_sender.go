package mocks

type FakeSender struct{}

func (s *FakeSender) Send(cmd, fprms, prms string) {}
