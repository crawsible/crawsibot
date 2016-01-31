package mocks

import "net"

type FakeDialer struct {
	DialCalls   int
	DialNetwork string
	DialAddress string

	DialReturnConn  *FakeConn
	DialReturnError error
}

func (d *FakeDialer) Dial(network, address string) (net.Conn, error) {
	d.DialCalls += 1
	d.DialNetwork = network
	d.DialAddress = address

	return d.DialReturnConn, d.DialReturnError
}
