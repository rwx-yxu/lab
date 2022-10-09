package urbase

import "fmt"

func Run() {
	fmt.Println("All bases")
	fmt.Println("Base: 2   8 10  16")
	for i := 0; i < 16; i++ {
		fmt.Printf(" %6b %3o %2d %3x\n", i, i, i, i)
	}
}
