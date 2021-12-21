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
		conn.Write([]byte("[+] Connection established\n[+] Message: "))
	}
	conn1, err := ln.Accept()
	if err != nil {
		conn1.Close()
		panic(err)
	} else if conn1 != nil {
		conn1.Write([]byte("[+] Connection established\n[-] You have to wait for the message from the other client\n"))
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		conn1.Write([]byte("\n[User 1]:"))
		conn1.Write([]byte(message))
		conn1.Write([]byte("\n[Type Message]:"))
		fmt.Printf("[+] Message: %s", message)

		if err != nil {
			conn.Close()
			conn1.Close()
			panic(err)
		}

		message1, err1 := bufio.NewReader(conn1).ReadString('\n')

		conn.Write([]byte("\n[User 2]:"))
		conn.Write([]byte(message1))
		conn.Write([]byte("\n[Type Message]:"))
		fmt.Printf("[+] Message: %s", message1)
		if err1 != nil {
			conn.Close()
			conn1.Close()
			panic(err1)
		}

	}
}
