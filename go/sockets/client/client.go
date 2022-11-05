package client

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Run(args []string) {
	service := args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println()
		os.Exit(1)
	}
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	result, err := io.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(result))
	os.Exit(0)
}
