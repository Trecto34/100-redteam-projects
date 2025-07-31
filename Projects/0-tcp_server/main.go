package main

import (
  "bufio"
  "fmt"
  "net"
)

func main() {
  port := ":1337"
  // Try to open the server for connections
  ln, err := net.Listen("tcp", port)
  if err != nil {
    panic(err)
  }
	
  fmt.Printf("[+] Server is running on port: %s\n", port )

  // Accept connection
  conn, err := ln.Accept()
  if err != nil {
    conn.Close()
    panic(err)
  }

  conn.Write([]byte("[+] Connection established\n"))

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
