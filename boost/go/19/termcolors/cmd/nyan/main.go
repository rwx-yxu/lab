package main

import (
	"fmt"
	"time"

	"github.com/rwx-yxu/lab/boost/go/19/termcolors"
)

func main() {
	for {
		fmt.Print(termcolors.Rand() + "nyan" + termcolors.Reset)
		time.Sleep(100 * time.Millisecond)
	}
}
