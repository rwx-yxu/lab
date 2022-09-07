package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string = "There"
	if len(os.Args[1:]) == 2 {
		name = strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("Hi, %v! \n", name)
}
