package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

func Throw(err any) {
	if err != nil {
		fmt.Println("error:", err)
		fmt.Printf("%s\n", debug.Stack())
		os.Exit(-1)
	}
}

var port = flag.Int("p", 80, "http listen port")

func main() {
	flag.Parse()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		host, err := os.Hostname()
		Throw(err)

		response := struct {
			Url           string
			Host          string
			ClientAddress string
			Hostname      string
			Now           int64
		}{
			Url:           ctx.Request.RequestURI,
			Host:          ctx.Request.Host,
			ClientAddress: ctx.ClientIP(),
			Hostname:      host,
			Now:           time.Now().Unix(),
		}
		ctx.JSON(200, &response)
	})
	r.Run(fmt.Sprintf(":%v", *port))
}
