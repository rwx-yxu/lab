# Boost week 19 notes
* The terminal that had won out for emulators is vt100
* vt100 is the most simple basics of terminal emulation. Every terminal
  will at least support it.
* vt100 did not support colors.
* The \$TERM env value follows ssh connection.
* To print colors follow (`ESC[color-codem`)
* `test -t 1` to test for non interactive terminals.
* Go will throw an error for `\e` when trying to declare to a string
  because it is an escape sequence.
    * Should use `\033` instead which is the octal notation.
* Using the `_test` naming convention for packages will help to catch
  export errors because Go will treat the test package as
  separate.
* Using %q in `fmt.Printf` will output the escape strings as
  hexadecimal. This is helpful for testing purposes for color output
  or when needing to check for an escape sequence.

```
//Use bash code in vi to print out code by !.bash
for i in {1..8}; do printf "const Color= \"\\\\033[3${i}m\"\n";
done
```

* Use ansi escape `\e[?25l` to hide cursor at start of program and
  `\e[?25h` to bring back cursor.
* Using `ctl + c` sends an interrupt signal to the command prompt.
* A handler is an function that runs when an event happens.
* Use trap in bash to capture signal signals.
  * `SIGTERM` is for a kill event.
  * `SIGINT` is for an interrupt.
  * Make sure to call trap again to avoid infinite trap loop.

```
//Bring back cursor if receieved an interrupt signal.
trap "printf \"\e[?25h\";exit; trap -- - SIGINT SIGTERM" SIGTERM SIGINT
```

* In Go, the math/rand package produces pseudo random numbers from a
  seed such that a seed will have a range of numbers to choose from.
  There is also another package crypto/rand which is less pseudo
  but it is costs more performance.

