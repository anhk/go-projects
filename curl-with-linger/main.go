package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/carlmjohnson/requests"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "invalid argument\n")
		os.Exit(-1)
	}

	u := args[1]

	if !strings.HasPrefix(u, "http://") && !strings.HasPrefix(u, "https://") {
		u = fmt.Sprintf("http://%v", u)
	}

	uu, err := url.Parse(u)
	Must(err)

	var netTransport = &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			conn, err := net.Dial(network, addr)
			if conn != nil && err == nil {
				err = conn.(*net.TCPConn).SetLinger(0)
			}
			return conn, err
		},
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	result := make(map[string]any)
	err = requests.URL(fmt.Sprintf("%s://%s", uu.Scheme, uu.Host)).
		Client(netClient).Path(uu.Path).ToJSON(&result).Fetch(context.Background())
	Must(err)

	data, err := json.MarshalIndent(result, "", "  ")
	Must(err)

	fmt.Printf("%s\n", data)
}

func Must(e any) {
	if e != nil {
		panic(e)
	}
}
