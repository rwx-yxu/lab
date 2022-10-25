package keeper

import (
	"fmt"
	"io"
	"strings"

	T "github.com/rwx-yxu/term"
	"github.com/rwx-yxu/term/sequence"
)

func Run(r io.Reader) {
	var name, colour string
	sequence.OnIfTerminal(r)
	//var err error
	fmt.Println(sequence.CLSEntire + sequence.Red + "Answer these questions three, the other side ye see." + sequence.Reset)
	fmt.Println(sequence.Red + "What is your name?" + sequence.Reset)
	name, _ = T.Prompt(r, sequence.Megenta+">"+sequence.Yellow)

	fmt.Println(sequence.Red + "What is your quest?" + sequence.Reset)
	T.Prompt(r, sequence.Megenta+">"+sequence.Yellow)

	if name == "Lancelot" || name == "Galahad" {
		fmt.Println(sequence.Red + "What is your favourite colour?" + sequence.Reset)
		colour, _ = T.Prompt(r, sequence.Megenta+">"+sequence.Yellow)
		if name == "Galahad" && strings.HasSuffix(colour, "no") {
			fmt.Println("You have perished")
		}
	}

	if name == "Robin" {
		fmt.Println(sequence.Red + "What is the capital of Assyria?" + sequence.Reset)
		T.Prompt(r, sequence.Megenta+">"+sequence.Yellow)
	}

	if name == "Arthur" {
		fmt.Println(sequence.Red + "What is the air speed velocity of and unladen swallow?" + sequence.Reset)
		T.Prompt(r, sequence.Megenta+">"+sequence.Yellow)
	}

}
