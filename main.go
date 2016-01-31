package main

import (
	"fmt"
	"net"
	"os"
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
}
