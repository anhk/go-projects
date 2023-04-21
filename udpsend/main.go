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

func main() {
	addr := flag.String("t", "", "udp server address")
	port := flag.String("p", "", "udp server port")
	data := flag.String("d", "test", "data to be sent")
	flag.Parse()
	if *addr == "" || *port == "" {
		flag.Usage()
		os.Exit(-1)
	}

	sock, err := net.Dial("udp", fmt.Sprintf("%v:%v", *addr, *port))
	Throw(err)

	_, err = sock.Write([]byte(*data))
	Throw(err)

	res := make([]byte, 8192)
	n, err := sock.Read(res)
	Throw(err)

	fmt.Println(string(res[:n]))
}
