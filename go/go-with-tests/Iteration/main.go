package main

import "fmt"

func Repeat(char string) {
	output := ""
	for i := 0; i < 6; i++ {
		output += char
	}

	fmt.Printf("%v", output)
}

func main() {
	Repeat("a")
}
