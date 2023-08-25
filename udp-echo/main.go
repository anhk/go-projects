package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"runtime/debug"
	"time"
)

func Throw(err any) {
	if err != nil {
		fmt.Println("error:", err)
		fmt.Printf("%s\n", debug.Stack())
		os.Exit(-1)
	}
}

var port = flag.Int("p", 33333, "udp listen port")

func main() {
	flag.Parse()

	sock, err := net.ListenUDP("udp", &net.UDPAddr{Port: *port})
	Throw(err)

	host, err := os.Hostname()
	Throw(err)

	data := make([]byte, 8192)

	response := struct {
		ClientAddress string
		Hostname      string
		Data          string
		Now           int64
	}{
		Hostname: host,
	}

	for {
		n, addr, err := sock.ReadFromUDP(data)
		Throw(err)

		response.ClientAddress = addr.IP.String()
		response.Data = string(data[:n])
		response.Now = time.Now().UnixNano()

		res, err := json.Marshal(response)
		Throw(err)

		sock.WriteToUDP(res, addr)
	}
}
