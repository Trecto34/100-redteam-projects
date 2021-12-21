package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Printf("[+] Server is running\n")
	// listen on port 1337
	ln, err := net.Listen("tcp", ":1337")
	if err != nil {
		panic(err)
	}

	// Accept connection
	conn, err := ln.Accept()
	if err != nil {
		conn.Close()
		panic(err)
	} else if conn != nil {
		conn.Write([]byte("[+] Connection established\n"))
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			conn.Close()
			panic(err)
		}

		fmt.Printf("[+] Message: %s", message)
		conn.Write([]byte("[+] Message received\n"))

	}
}
