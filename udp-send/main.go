package main

import (
	"flag"
	"fmt"
	"go-projects/pkg/log"
	"net"
	"os"
	"runtime/debug"
	"strings"
)

func Throw(err any) {
	if err != nil {
		fmt.Println("error:", err)
		fmt.Printf("%s\n", debug.Stack())
		os.Exit(-1)
	}
}

var (
	addr    = flag.String("t", "", "udp server address")
	port    = flag.Int("p", 33333, "udp server port")
	data    = flag.String("d", "test", "data to be sent")
	typ     = flag.Bool("1", false, "use sendmsg instead of connect")
	datalen = flag.Int("s", 4, "data size")
)

func main() {
	flag.Parse()

	if *addr == "" {
		flag.Usage()
		os.Exit(-1)
	}

	if *typ {
		SendMessageBySendMsg()
	} else {
		SendMessageByConnect()
	}
}

func expandToDatalen(data string, datalen int) []byte {
	var arr []string
	for {
		if datalen > len(data) {
			arr = append(arr, data)
			datalen -= len(data)
		} else {
			arr = append(arr, data[:datalen])
			break
		}
	}

	return []byte(strings.Join(arr, ""))
}

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func SendMessageBySendMsg() {
	localPort := RandomInt(40000, 50000)
	remotePort := *port

	sock, err := net.ListenUDP("udp", &net.UDPAddr{Port: localPort})
	Throw(err)

	remoteAddr := net.UDPAddr{IP: net.ParseIP(*addr), Port: remotePort}

	log.Debugf("[sendmsg] udp remote address: %v", remoteAddr.String())

	_, err = sock.WriteTo(expandToDatalen(*data, *datalen), &remoteAddr)
	Throw(err)

	res := make([]byte, 8192)
	n, _, err := sock.ReadFromUDP(res)
	Throw(err)

	fmt.Println(string(res[:n]))
}

func SendMessageByConnect() {
	remoteAddr := net.UDPAddr{IP: net.ParseIP(*addr), Port: *port}
	log.Debugf("[connect] udp remote address: %v", remoteAddr.String())

	sock, err := net.Dial("udp", remoteAddr.String())
	Throw(err)

	_, err = sock.Write(expandToDatalen(*data, *datalen))
	Throw(err)

	res := make([]byte, 8192)
	n, err := sock.Read(res)
	Throw(err)

	fmt.Println(string(res[:n]))
}
