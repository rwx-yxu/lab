package main

import (
	"fmt"
	"log"

	"github.com/rwx-yxu/boost/hello/foo"
)

func main() {
	println("Hello world. println")         // os.Stderr
	fmt.Println("Hello world. fmt.Println") // os.Stdout
	log.Println("Hello world. log.Println") //os.Stderr
	foo.Hello()
}
