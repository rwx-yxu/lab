package hello_test

import "testing"

func TestPrintHelloMessageToTerminal(t *testing.T) {
	t.Parallel()
	hello.Print()
}
