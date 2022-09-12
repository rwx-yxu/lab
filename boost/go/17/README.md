# Boost week 17 notes
* Redo the hello world challenge in go.
* Redirect standard out - go run main.go 2>err 1>out. Redirects
  the outputs to file called err and another file called out.
* `println` and `log.Println` prints to standard error where as
  `fmt.Println` prints to standard output.
* Always make a distinction with output. If the output is usage
  information or error information that does not relate to what the
  program is actually doing, always log it.
* `log.Println` is better because the output can be redirected to a
  centralised log compared to `fmt.Println`
* `println` is not official part of go.
* Interpreted languages like python and bash ect, have an "implicit" main
* A go module is usually associated with a single git repo. Where as a
  package are sub directories of a module.
* Go run only works on files that have no external dependencies on
  anythings else.
* Go will support multiple modules in the same repo. 
* `entr bash -c "clear; go run ." < <(find .)` to run go test
  continuously for any change to a file in current directory.
* The output line for example tests in go compares the output comment
  with standard out. This means that using `log.Println` will not be
  used for the output for example tests.
* It is common for main just to call another
  function so that testing can be easier.
* During go test, the standard error output is separated from the
  results tests. This means that you can see debug information
  separate from the actual result when developing. This works well
  because using `log.println` doesn't impact the standard output.
* `os.Getpid()` gets the current process id.
* Generally default to export everything and use internal package when
  needed. This makes it easier to move code around later.
* Internal package is only visible to the parent directory.
* Use `shift+k` in vi when the cursor is over a function to bring up
  the documentation. Exit with `:q`
* Remember that for example tests, the comment `//Output:` requires the
  colon for go test to pick up.
* For example testing you can add an `_` after the function name for
  additional context. The function name must match in order for go
  reflection to find it.
* Adding ... to a variable type in a function declaration will make the
  parameter optional
* Cannot test main function because go test generates a temp test binary
  which throws off our argument count.
