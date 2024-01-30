package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 65002})
	check(err)

	var peers = make([]*net.UDPAddr, 0, 2)
	var data = make([]byte, 8192)

	for {
		n, remote, err := listener.ReadFromUDP(data)
		check(err)

		fmt.Printf("read from %v: %v\n", remote.String(), string(data[:n]))
		peers = append(peers, remote)

		if len(peers) == 2 {
			listener.WriteToUDP([]byte(peers[0].String()), peers[1])
			listener.WriteToUDP([]byte(peers[1].String()), peers[2])
			break
		}
	}
	time.Sleep(time.Second * 5)
}

func check(e any) {
	if e != nil {
		panic(e)
	}
}
