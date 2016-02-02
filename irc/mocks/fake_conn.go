package mocks

import (
	"net"
	"time"
)

type FakeConn struct{}

func (c *FakeConn) Read(b []byte) (int, error)         { return 0, nil }
func (c *FakeConn) Write([]byte) (int, error)          { return 0, nil }
func (c *FakeConn) Close() error                       { return nil }
func (c *FakeConn) LocalAddr() net.Addr                { return &FakeAddr{} }
func (c *FakeConn) RemoteAddr() net.Addr               { return &FakeAddr{} }
func (c *FakeConn) SetDeadline(t time.Time) error      { return nil }
func (c *FakeConn) SetReadDeadline(t time.Time) error  { return nil }
func (c *FakeConn) SetWriteDeadline(t time.Time) error { return nil }

type FakeAddr struct{}

func (a *FakeAddr) Network() string { return "" }
func (a *FakeAddr) String() string  { return "" }
