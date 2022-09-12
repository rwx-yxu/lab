package main

import (
	"fmt"
	"os"
	"strings"
)

func SayHi(args ...string) {
	name := "there"
	if len(args) > 0 {
		name = strings.Join(args, " ")
	}
	fmt.Printf("Hi %v!\n", name)
}

func main() {
	SayHi(os.Args[1:]...)
}
