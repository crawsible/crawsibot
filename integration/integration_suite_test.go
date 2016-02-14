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
	reqCh = make(chan string, 1000)
	resCh = make(chan string, 1000)

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
	var conn net.Conn
	var err error

	go func() {
		for msg := range resCh {
			fmt.Fprintf(conn, msg)
		}
	}()

	for {
		conn, err = ln.Accept()
		if err != nil {
			panic(err.Error())
		}

		reader := bufio.NewReader(conn)
		var line string

		for {
			line, err = reader.ReadString('\n')
			if err != nil {
				break
			}

			reqCh <- line
		}
	}
}
