package tcpserver

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/rwx-yxu/term"
)

func Run() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//net.Listen makes the program a TCP server. Returns a Listener
	//variable
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer listener.Close()
	//Only after a successful call to Accept the TCP server can begin to
	//interact with TCP clients.
	//Since Accept() is outside the for loop, this TCP server can only
	//serve one client.
	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {

		netData, _ := term.ReadLine(conn)

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Printf("-> %s\n", string(netData))

		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		conn.Write([]byte(myTime))
	}
}
