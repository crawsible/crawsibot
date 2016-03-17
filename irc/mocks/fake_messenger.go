package mocks

type FakeMessenger struct {
	EnrollForMsgsCalls int
	EnrollForMsgsRcvr  chan map[string]string
	EnrollForMsgsCmd   string

	SendArgs  [][]string
	sendCalls int
}

func (m *FakeMessenger) EnrollForMsgs(rcvCh chan map[string]string, cmd string) {
	m.EnrollForMsgsCalls += 1
	m.EnrollForMsgsRcvr = rcvCh
	m.EnrollForMsgsCmd = cmd
}

func (m *FakeMessenger) Send(cmd, fprms, prms string) {
	m.sendCalls += 1
	m.SendArgs = append(m.SendArgs, []string{cmd, fprms, prms})
}

func (m *FakeMessenger) SendCalls() int {
	return m.sendCalls
}
