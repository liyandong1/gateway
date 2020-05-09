package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
		Zone: "",
	})
	if err != nil {
		fmt.Printf("connect failed, err: %v\n", err)
		return
	}

	for i := 0; i < 100; i++ {
		_, err := conn.Write([]byte("Hello server!"))
		if err != nil {
			fmt.Printf("send data failed, err: %v\n", err)
			return
		}

		result := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(result)
		if err != nil {
			fmt.Printf("received data failed, err: %v", err)
			return
		}
		fmt.Printf("receive from addr: %v data: %v\n", remoteAddr, string(result[:n]))
	}
}
