package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		fmt.Printf("Listen failed, err:%v\n", err)
		return
	}

	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("Read failed from addr: %v, err: %v\n", addr, err)
			break
		}

		go func() {
			fmt.Printf("addr: %v data: %v  count: %v\n", string(data[:n]), n)
			_, err = listen.WriteToUDP([]byte("received success!!!"), addr)
			if err != nil {
				fmt.Printf("Write failed, err: %v\n", err)
			}
		}()
	}
}
