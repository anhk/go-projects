package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime/debug"
)

func Throw(err any) {
	if err != nil {
		fmt.Println("error:", err)
		fmt.Printf("%s\n", debug.Stack())
		os.Exit(-1)
	}
}

var (
	addr = flag.String("t", "", "udp server address")
	port = flag.Int("p", 33333, "udp server port")
	data = flag.String("d", "test", "data to be sent")
	typ  = flag.Bool("1", false, "use sendmsg instead of connect")
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

	remoteAddr := net.UDPAddr{
		IP:   net.ParseIP(*addr),
		Port: remotePort,
	}

	_, err = sock.WriteTo([]byte(*data), &remoteAddr)
	Throw(err)

	res := make([]byte, 8192)
	n, _, err := sock.ReadFromUDP(res)
	Throw(err)

	fmt.Println(string(res[:n]))
}

func SendMessageByConnect() {

	sock, err := net.Dial("udp", fmt.Sprintf("%v:%v", *addr, *port))
	Throw(err)

	_, err = sock.Write([]byte(*data))
	Throw(err)

	res := make([]byte, 8192)
	n, err := sock.Read(res)
	Throw(err)

	fmt.Println(string(res[:n]))
}
