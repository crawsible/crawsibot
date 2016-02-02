package mocks

type FakeWriter struct {
	WriteMessage []byte
	writeCalls   int
}

func (c *FakeWriter) Write(b []byte) (int, error) {
	c.WriteMessage = b
	c.writeCalls += 1
	return 0, nil
}

func (c *FakeWriter) WriteCalls() int {
	return c.writeCalls
}
