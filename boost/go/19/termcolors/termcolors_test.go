package termcolors_test

import (
	"fmt"

	"github.com/rwx-yxu/lab/boost/go/19/termcolors"
)

func ExampleBlack() {
	fmt.Printf("%q\n", termcolors.Black+"black"+termcolors.Reset)
	fmt.Printf("%q\n", termcolors.Red+"red"+termcolors.Reset)
	fmt.Printf("%q\n", termcolors.Green+"green"+termcolors.Reset)
	fmt.Printf("%q\n", termcolors.Yellow+"yellow"+termcolors.Reset)
	fmt.Printf("%q\n", termcolors.Blue+"blue"+termcolors.Reset)
	fmt.Printf("%q\n", termcolors.Megenta+"magenta"+termcolors.Reset)
	fmt.Printf("%q\n", termcolors.Cyan+"cyan"+termcolors.Reset)
	fmt.Printf("%q\n", termcolors.White+"white"+termcolors.Reset)

	//Output:
	//"\x1b[30mblack\x1b[0m"
	//"\x1b[31mred\x1b[0m"
	//"\x1b[32mgreen\x1b[0m"
	//"\x1b[33myellow\x1b[0m"
	//"\x1b[34mblue\x1b[0m"
	//"\x1b[35mmagenta\x1b[0m"
	//"\x1b[36mcyan\x1b[0m"
	//"\x1b[37mwhite\x1b[0m"
}

func ExampleRand() {
	fmt.Printf("%q", termcolors.Rand())
}
