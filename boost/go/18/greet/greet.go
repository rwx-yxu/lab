package main

import (
	"fmt"
	"os"
	"strings"
)

func greet(args ...string) {
	name := "there"
	if len(args) > 0 {
		name = strings.Join(args, " ")
	}
	fmt.Printf("Hi, %s!\n", name)
}

func main() {
	greet(os.Args...)
}
