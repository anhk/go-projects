package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime/debug"
	"strconv"
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
	port = flag.String("p", "", "udp server port")
	data = flag.String("d", "test", "data to be sent")
	typ  = flag.Bool("1", false, "use sendmsg instead of connect")
	v6   = flag.Bool("6", false, "use ipv6 network")
)

func main() {
	flag.Parse()

	if *addr == "" || *port == "" {
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
	remotePort, err := strconv.Atoi(*port)
	Throw(err)

	proto := If(!*v6, "udp4", "udp")

	sock, err := net.ListenUDP(proto, &net.UDPAddr{Port: localPort})
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
