package internal_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/rwx-yxu/lab/boost/go/18/greet/internal"
)

func ExampleReadLine() {
	//Turn string into standard input
	//As if it is standard input
	sr := strings.NewReader("Sample\r\n")
	line, err := internal.ReadLine(sr)
	if err != nil {
		log.Println(err)
		return
	}
	//Using %s will pass
	fmt.Printf("%q", line)
	//Output:
	//"Sample"

}
