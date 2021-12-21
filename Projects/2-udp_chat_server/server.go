package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 1337,
		IP:   []byte{0, 0, 0, 0},
		Zone: "",
	})
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("[+] Server is running on %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, 128)
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}
		conn.WriteToUDP([]byte("[+] Message: "), remote)
		data := strings.TrimSpace(string(message[:rlen]))
		fmt.Printf("[+] Message: [%s] from (%s)\n", data, remote)
	}
}
