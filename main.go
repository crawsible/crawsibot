package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Printf("Danger, Will Robinson! %s", err.Error)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "PASS some-password\r\nNICK some-username\r\n")
	fmt.Fprintf(conn, "JOIN #crawsible\r\n")
	fmt.Fprintf(conn, "PRIVMSG #crawsible :I HAVE ARRIVED.\r\n")
	ircReader := bufio.NewReader(conn)

	ch := make(chan string)
	eCh := make(chan error)

	go func(ch chan string, eCh chan error) {
		var line string
		var err error
		for {
			line, err = ircReader.ReadString('\n')
			if err != nil {
				eCh <- err
				break
			}
			ch <- line
		}
	}(ch, eCh)

	for {
		select {
		case line := <-ch:
			fmt.Printf("%s", line)
			if match, _ := regexp.MatchString("PING :tmi.twitch.tv\r?\n", line); match {
				fmt.Fprintln(conn, "PONG :tmi.twitch.tv\n")
				fmt.Println("Responding: \"PONG :tmi.twitch.tv\"")
			}
		case err := <-eCh:
			fmt.Printf("ircReader error: ", err.Error)
			break
		}
	}
}
