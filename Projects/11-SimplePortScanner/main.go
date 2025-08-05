package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func scanner(host string){
	dialer := net.Dialer{
		Timeout: 500 * time.Millisecond,
	}

	for port := 1; port < 1024; port++  {
		porta := strconv.Itoa(port)
		address := net.JoinHostPort(host, porta)
		conn, err := dialer.Dial("tcp", address)
		if err != nil {
			continue
		}
		defer conn.Close()
		fmt.Printf("Success connecting to %s on port %s\n", host, porta)
	}

}

func main()  {
	host := "scanme.nmap.org"
	go scanner(host)
}
