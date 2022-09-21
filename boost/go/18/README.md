# Boost week 18 notes
* Use `go install ./cmd/greet/; greet` in go mod directory to install the binary to local
  machine from the go env variable.
* The cmd directory is for indevidual commands
* Stuff that has to do with interacting with the user on the command
  line, should not be in the package library.

  * The package library should be independent no matter
    what. We should not include details on how to obtain the
    parameter in the package library itself. Since that is
    stuff is a detail about how the command executes.
* Interactive terminal in general is a bad idea for a production
  environment. It is better to prefer arguments, env vars or
  configuration files.
  * This can cause stalling by waiting for an input.
  * For example, not using -y for apt-get.
* Unix filters are designed to read standard input,
  modify and produce an output. This is a use case where by you expect
  the program to block and look for standard input.
* Internal directory cannot be accessed by other modules.
* Always name test packages with `_test` for example tests. This
  would simulate an example as if another user were to use the
 test package.
* Can also use `Unordered Output:` test for examples. This would mean
  that the line order doesn't matter so long as the outputs match.
* The go sandbox saves the go documentation whereby you can run the
  code. Using a test framework like testify will not allow for this
  though.
* Use `fmt.Printf("%q")` to print the string literal. Basically, see
  the invisible characters like `\n` or `\r`. Useful to properly check
  for an empty value.
* Convention for go documentation is to always start with the name.
* Use `go doc -all .` to check documented code with comments.
* Common to create a file called `doc.go`.
