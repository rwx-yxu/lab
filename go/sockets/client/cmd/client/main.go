package main

import (
	"fmt"
	"os"

	"github.com/rwx-yxu/lab/go/sockets/client"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	client.Run(os.Args, os.Stdin)
}
