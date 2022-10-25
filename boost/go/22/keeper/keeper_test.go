package keeper_test

import (
	"strings"

	"github.com/rwx-yxu/lab/boost/go/22/keeper"
)

func ExampleRun() {
	// normallu pass os.Stdin (this is just for testing)
	//Colours are disabled unless interactive terminal detected
	r := strings.NewReader("Galahad\n")
	keeper.Run(r)
	//Output:
	//Answer these questions three, the other side ye see.
	//What is your name?
	//>What is your quest?
	//>What is your favourite colour?
	//>
}
