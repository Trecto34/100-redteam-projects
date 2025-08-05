package main

import (
  "bufio"
  "fmt"
  "net"
)

func main(){
	Server()
}

func Encrypt(alpha string) string {
	var encryptedMessage string
	for _, letter := range alpha {
		encryptedMessage += string(Shift(30, letter))
	}
	return encryptedMessage
}

func Shift(pace int, lett rune) byte {
	letter := byte(lett) + byte(pace)
	for letter > 90 {
		letter = letter - 26
	}
	return letter
}

func Server() {
	// listen on port 1337
	ln, err := net.Listen("tcp", ":1337")
	if err != nil {
		panic(err)
	}
	fmt.Printf("[+] Server is running\n")

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
		go client(conn1, conn)
		go client(conn, conn1)
	}
}

func client(connection net.Conn, other_client net.Conn) {
	message, err := bufio.NewReader(other_client).ReadString('\n')

	encryptedMessage := Encrypt(message)
	connection.Write([]byte("\n[User]: "))
	connection.Write([]byte(encryptedMessage))
	connection.Write([]byte("\n[Type Message]: "))
	fmt.Printf("[+] Encrypted Message: %s", encryptedMessage)

	if err != nil {
		connection.Close()
		other_client.Close()
		panic(err)
	}
}
