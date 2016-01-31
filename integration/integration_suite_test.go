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

func TestIntegration(t *testing.T) {
	reqCh = make(chan string)
	resCh = make(chan string)

	RegisterFailHandler(Fail)
	go mockIRCServer(reqCh, resCh)

	RunSpecs(t, "Integration Suite")
}

func mockIRCServer(reqCh, resCh chan string) {
	ln, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		panic(err.Error())
	}
	defer ln.Close()

	serverCycle(ln, reqCh, resCh)
}

func serverCycle(ln net.Listener, reqCh, resCh chan string) {
	conn, err := ln.Accept()
	if err != nil {
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	var line string

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if line == "EOF\r\n" {
			fmt.Println("signing off!")
			return
		}

		reqCh <- line
	}
}
