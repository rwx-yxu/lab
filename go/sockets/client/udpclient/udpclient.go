package udpclient

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/rwx-yxu/term"
)

func Run(args []string, r io.Reader) {
	s, err := net.ResolveUDPAddr("udp4", args[1])
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())

	defer c.Close()

	for {
		req, _ := term.Prompt(r, ">>")

		data := []byte(req + "\n")

		_, err = c.Write(data)

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

		if err != nil {
			log.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		//Reads a packet from the server connection and will return if an
		//error occurs
		n, _, err := c.ReadFromUDP(buffer)

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}

}
