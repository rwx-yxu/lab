package greet_test

import "github.com/rwx-yxu/lab/boost/go/18/greet"

/*
func ExampleGreet() {

	greet.Hi()
	//Output:
	//Hi, there!
}
*/
func ExampleGreet_With_Arguments() {
	greet.Hi("Yongle")
	//Output:
	//Hi, Yongle, nice to meet you!
}
