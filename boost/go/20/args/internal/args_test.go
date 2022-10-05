package internal_test

import (
	"fmt"

	"github.com/rwx-yxu/lab/boost/go/20/args/internal"
)

/*
const testone = `
$0 --> "./args"
$1 --> "first"
$2 --> "second"
$3 --> "third"
`

func TestOutput(t *testing.T) {
	out := output()
	if out != testone {
		t.Errorf("\nwant: %q\ngot:  %q\n", testone, out)
	}
}
*/

func ExampleOutput() {
	fmt.Println(internal.Output("./args", "first", "second", "third"))
	//Output:
	//$0 --> "./args"
	//$1 --> "first"
	//$2 --> "second"
	//$3 --> "third"
}
