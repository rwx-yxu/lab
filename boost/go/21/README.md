# Week 21 boost notes
* Example tests will always cut out white space for the first output.
* Place a number after `%` in `fmt.Printf` to indent to the right by n
  spaces.
* Go work allows for multi package development without having to pull
  down the repository from the internet.
* Used my terminal package for colors when making the waffles
  challenge.
* Proper way to install anything go from the command line is to use Go
  install@latest or version. **Not** go get. The purpose of go get is to
  pull down a dependency and to freeze it at that version for the go
  module to use. Never has been for installing commands.
* No need to reference the imported module name when calling from
  package of imported module.
* Replace in go.mod should only be used if a dependency changes the
  repository name. So we do not have to rename everything.
* Using replace incorrectly to act like go work will lock projects to
  that version of go unable to go to 1.18 because it will stop
  compilation for 1.18 and above. Use go work instead.
* A common convention in go is when a function begins with `Must`,
  instead of an error being returned, the function will panic instead.
  This means that the code will not even compile if any failure happens
  in a `Must` function.
* A byte slice is different to a string. You can always cast a byte
  slice into a string and vice versa.
* Regex is expensive to run. A regular substring matching will be more
  performant.
* Any regular expression has to be compiled.
  * In an interpreted language, the regex has to be compiled and then
    ran every time.
  * This is why interpreted languages are slower than languages that can
    pre-compile the regular expression.
  * Do not put regex compile expressions in functions that will be
    called multiple times in Go! This will cause the regex to be
    compiled every time the function is called. Instead, place the regex
    as a global variable so it is only compiled once at start up.
* Most times, avoid regex and use substrings instead.
* It is idiomatic go to have `execptions` to be indented. Your main
  concerns are to the left.
* Use var for zero value declaration of variables.
* Garbage collection is how the language deals with memory utilization.
* The go.sum file locks down a dependency version at the time it was
  pulled down. This is updated whenever another go get is done.
