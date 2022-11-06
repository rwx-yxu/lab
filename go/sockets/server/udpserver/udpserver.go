package udpserver

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func Run(args []string) {
	s, err := net.ResolveUDPAddr("udp4", args[1])
	if err != nil {
		log.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp4", s)

	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		fmt.Printf("->%s \n", string(buffer[0:n-1]))
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err = conn.WriteToUDP(data, addr)
		if err != nil {
			log.Println(err)
			return
		}

	}

}

func random(min, max int) int {

	return rand.Intn(max-min) + min
}
