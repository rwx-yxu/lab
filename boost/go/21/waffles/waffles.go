package waffles

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/rwx-yxu/term"
	S "github.com/rwx-yxu/term/sequence"
)

func prompt(in io.Reader) string {
	resp, err := term.Prompt(in, S.Megenta+"> "+S.Yellow)
	if err != nil {
		log.Println(err)
		return ""
	}
	fmt.Print(S.Reset)

	return resp

}

func end() {
	fmt.Println("The end")
	os.Exit(1)
}

var YES = regexp.MustCompile(`(?i)^(y(es|eah|ep)|sure|ok)`)

func Run(in io.Reader) {
	var resp string
	fmt.Println(S.Red + "Do you like waffles?" + S.Yellow)
	resp = prompt(in)
	if !YES.MatchString(resp) {
		end()
	}

	fmt.Println(S.Red + "Do you like pancakes?" + S.Yellow)
	resp = prompt(in)
	if !YES.MatchString(resp) {
		end()
	}

}
