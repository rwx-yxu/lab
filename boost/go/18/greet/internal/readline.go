package internal

import (
	"bufio"
	"io"
	"strings"
)

//Readline takes any io.Reader and returns a trimmed string (initial and trailing white space) or an empty
//string if an error is encountered
func ReadLine(in io.Reader) (string, error) {
	//TODO Consider using a scanner
	//This is technically wrong since it will read    until the line
	//return. There might be a carriage return before that.
	out, err := bufio.NewReader(in).ReadString('\n')
	out = strings.TrimSpace(out)
	return out, err
}
