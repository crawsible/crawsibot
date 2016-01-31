package integration_test

import (
	"bufio"
	"fmt"
	"net"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var reqCh chan string
var resCh chan string
var quitCh chan struct{}

func TestIntegration(t *testing.T) {
	reqCh = make(chan string)
	resCh = make(chan string)
	quitCh = make(chan struct{})

	RegisterFailHandler(Fail)
	go mockIRCServer(reqCh, resCh, quitCh)

	RunSpecs(t, "Integration Suite")
	fmt.Println("TESTS COMPLETE")
	quitCh <- struct{}{}
}

func mockIRCServer(reqCh, resCh chan string, quitCh chan struct{}) {
	ln, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		panic(err.Error())
	}
	defer ln.Close()

	for {
		select {
		case <-quitCh:
			close(reqCh)
			close(resCh)
			close(quitCh)

			return
		default:
			serverCycle(ln, reqCh, resCh)
		}
	}
}

func serverCycle(ln net.Listener, reqCh, resCh chan string) {
	conn, err := ln.Accept()
	if err != nil {
		return
	}

	reader := bufio.NewReader(conn)
	go func(reqCh, resCh chan string) {
		defer conn.Close()
		var (
			line string
			err  error
		)

		for {
			line, err = reader.ReadString('\n')
			if err != nil {
				panic(err.Error())
			}

			reqCh <- line
		}
	}(reqCh, resCh)
}
