package eightball_test

import (
	"fmt"

	"github.com/rwx-yxu/lab/boost/go/20/eightball"
)

func ExampleRespond() {
	resp := eightball.Respond()
	fmt.Println(len(resp) > 0)
	//Output:
	//true
}
