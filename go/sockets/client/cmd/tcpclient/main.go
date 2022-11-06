package main

import (
	"fmt"
	"os"

	"github.com/rwx-yxu/lab/go/sockets/client/tcpclient"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	tcpclient.Run(os.Args, os.Stdin)
}
