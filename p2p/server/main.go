package main

import (
	"bytes"
	"encoding/binary"
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
			r1 := &bytes.Buffer{}
			binary.Write(r1, binary.BigEndian, peers[0])
			listener.WriteToUDP(r1.Bytes(), peers[1])

			r2 := &bytes.Buffer{}
			binary.Write(r2, binary.BigEndian, peers[1])
			listener.WriteToUDP(r2.Bytes(), peers[0])
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
