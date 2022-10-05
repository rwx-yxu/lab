package main

import (
	"fmt"
	"os"

	"github.com/rwx-yxu/lab/boost/go/20/args/internal"
)

func main() {
	fmt.Print(internal.Output(os.Args...))
}
