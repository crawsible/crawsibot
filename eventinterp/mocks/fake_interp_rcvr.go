package mocks

type FakeInterpRcvr struct {
	loggedInCalls int
}

func (r *FakeInterpRcvr) LoggedIn() {
	r.loggedInCalls += 1
}

func (r *FakeInterpRcvr) LoggedInCalls() int {
	return r.loggedInCalls
}
