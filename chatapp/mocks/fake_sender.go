package mocks

type FakeSender struct {
	SendCmd   string
	SendFprms string
	SendPrms  string

	sendCalls int
}

func (s *FakeSender) Send(cmd, fprms, prms string) {
	s.sendCalls += 1
	s.SendCmd = cmd
	s.SendFprms = fprms
	s.SendPrms = prms
}

func (s *FakeSender) SendCalls() int {
	return s.sendCalls
}
