package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var server string

func runAsClient() {
	local := &net.UDPAddr{IP: net.IPv4zero, Port: 65001}
	dest := &net.UDPAddr{IP: net.ParseIP(server), Port: 65002}

	conn, err := net.DialUDP("udp", local, dest)
	check(err)

	_, err = conn.Write([]byte("hello world."))
	check(err)

	var data = make([]byte, 8192)
	n, _, err := conn.ReadFromUDP(data)
	check(err)
	fmt.Printf("read from remote %d\n", n)

	peer, err := net.ResolveUDPAddr("udp", string(data[:n]))
	check(err)

	binary.Read(bytes.NewReader(data[:n]), binary.BigEndian, &peer)
	fmt.Printf("peer: %v", peer.String())
	conn.Close()

	conn, err = net.DialUDP("udp", local, peer)
	check(err)

	go func() {
		for i := 0; i < 10; i++ {
			_, err := conn.Write([]byte(fmt.Sprintf("-- %d -- ", i)))
			check(err)
			time.Sleep(time.Second)
		}
	}()

	for {
		n, r, err := conn.ReadFromUDP(data)
		check(err)

		fmt.Printf("read from %v: %v\n", r.String(), string(data[:n]))
	}

}

var rootCmd = cobra.Command{
	Use: "",
	Run: func(cmd *cobra.Command, args []string) {
		runAsClient()
	},
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "127.0.0.1", "server address")
	rootCmd.Execute()
}

func check(e any) {
	if e != nil {
		panic(e)
	}
}
