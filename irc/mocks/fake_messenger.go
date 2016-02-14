package mocks

import "github.com/crawsible/crawsibot/irc"

type FakeMessenger struct {
	EnrollForPINGCalls     int
	EnrollForPINGRecipient irc.PINGRecipient

	SendArgs  [][]string
	sendCalls int
}

func (m *FakeMessenger) EnrollForPING(rcp irc.PINGRecipient) {
	m.EnrollForPINGCalls += 1
	m.EnrollForPINGRecipient = rcp
}

func (m *FakeMessenger) Send(cmd, fprms, prms string) {
	m.sendCalls += 1
	m.SendArgs = append(m.SendArgs, []string{cmd, fprms, prms})
}

func (m *FakeMessenger) SendCalls() int {
	return m.sendCalls
}
