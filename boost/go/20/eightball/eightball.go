package eightball

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/rwx-yxu/lab/boost/go/20/cli"
)

const art = `
        ____
    ,dP9CGG88@b,
  ,IP  _   Y888@@b,
 dIi  (_)   G8888@b
dCII  (_)   G8893@@b
GCCIi     ,GG8888@@@
GGCCCCCCCGGG88888@@@
GGGGCCCGGGG88888@@@@...
Y8GGGGGG8888888@@@@P.....
 Y88888888888@@@@@P......
 Y8888888@@@@@@@P'......
    @@@@@@@@@P'.......
        """"........

`

//Using [...] ensures that it is an array and not slice
var Responses = [...]string{
	"Yes",
	"No",
	"Maybe",
	"Never",
}

//Respond will return a random response from a list of Responses.
func Respond() string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(Responses))

	//Pick response
	return Responses[i]
}

//Run starts an interactive eight ball session prompting the user for
//input and answering until interrupted.
func Run() {
	fmt.Println(art)
	fmt.Println("Welcome to the magic eight ball!")
	fmt.Println("Enter your yes or no question")
	for {
		cli.Prompt(os.Stdin, "-->")
		fmt.Println(Respond())
	}
}
