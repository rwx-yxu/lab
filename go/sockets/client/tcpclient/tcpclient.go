package tcpclient

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/rwx-yxu/term"
)

func Run(args []string, r io.Reader) {
	service := args[1]
	//Verify TCP address:port first
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//net.DialTCP begins the implementations of the TCP client to desired
	//TCP server.
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {
		req, _ := term.Prompt(r, ">>")

		//Write to connection
		fmt.Fprintf(conn, req+"\n")

		//Read response
		resp, _ := term.ReadLine(conn)
		fmt.Printf("Resp:%s\n", resp)

		//Only terminate when receive request "STOP"
		if strings.TrimSpace(string(req)) == "STOP" {
			fmt.Println("TCP Client exiting...")
			return
		}
	}
}
