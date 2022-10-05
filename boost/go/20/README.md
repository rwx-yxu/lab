# Boost week 20
* Main function should always be really simple because of how difficult it
is to test.
* Do not import testing package when using example test. The example test
will not be evaluated if there is a failure of a higher test using the
testing package.
* Use `fmt.Printf("%T")` to print the variable type
* The go test command packs in additional args for testing. This makes
testing a program with args tricky when you want to check the number of
args.
* A programming paradigm is a style, technique or way of writing a
  program.
* The three main paradigms are: procedural, functional and object
  oriented.
* Go is a hybrid language. Same with python. Java is a very strict
  object oriented program.
* Difference between a function and procedure is that all functions are
  procedures but not all procedures are functions.
* Functional has no side effects.
* Functional programming is easier to test because functions do not care
  about anything else.

```
//Following code not testable because there is no parameter being passed
into the output function.
//The use of os.Args is not testable when trying to make a test
function because os.Args is pulled outside of the scope of the function.
//Should pass in an array for args as a parameter to func output.
func output() string {
  var buf string
  for n, val := range os.Args {
    buf += fmt.Sprintf("$%v --> %q\n", n, val)
  }
  return buf
}
```

* Sometimes, just running the program is better than writing out a unit
  test.
* Use exported symbols by default.
* Use internal package to make packages visible to current module but
  prevent external modules to use. This prevents packages from being
  from the internal package if imported else where.
* The ... in go means that it makes the variable into a slice.
* Array vs slice? An array is tied to memory which means it cannot
  change where as a slice is dynamic and can change. Slices are always
  pointers. Should mostly use slices.
* Slices are an abstraction on top of arrays in Go.
* An edge case of a slice is that if you have a large slice and only
  want a range of values from that slice, the underlying slice
  referenced will still be kept alive. To avoid this, create a new slice
  and populate with the value ranges you want from the larger slice and
  then dereference to free up memory.
* `Getops` is very prevalent in the many command line applications but
  it is not user friendly because it is not standardized. Sometimes, you
  need to prepend with a `--` or `-` or use short notation like `-h` and
  `-help`. Even then, you might or might not need the `=` when setting a
  value.
* Go requires random to be reseeded very time or else the same number
  will be picked.
* Try to avoid naming modules `util`
* Go is only aware of modules that has been committed to the internet.
* Use `go work init` to resolve local dependencies that are not on the
  internet or require to use a local version.
* Usually, you do not want to commit the go.work to source control
  though. Will cause issues if some else tries to build but they do not
  have the dependency in a specific place defined in the go.work file.

