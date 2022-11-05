package tcpserver

import (
	"log"
	"net"
	"os"
)

func Run() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		resp := "Hello world"
		conn.Write([]byte(resp))
		//Close connection
		conn.Close()
	}
}
