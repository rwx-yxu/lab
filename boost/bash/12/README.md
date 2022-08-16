# Boost week 12 notes
* Alternative terms for `array`
  * `lists`
  * `vector`
  * `tuple` 
  * `arrayList`
* cat <<'EOM' command prints out exactly the text after.
  * This is seen a lot in kubernetes as to pass through yaml to
    kubctl . The text comes into standard input so echo doesn't
    work.
  * :`set noet` to set tabs. `set nolist` to show the spacing. Using
    tabs should use `cat <<-EOM` instead. Otherwise, this command
    will not work.
* echo adds a `/n` at the end of the line. This can be an issue if you
  want to read an input on the same line. `printf` can be used instead.
* `while true; do clear; date; sleep .5; done ` creates an infinite
  loop.
* man pages has sections. Using help will fixate on the bash man page
  only. This is useful when looking for commands that specifically
  relate to bash.
* Arrays in bash do not require a comma to separate out the items when
  declaring.
* `printf` will execute once per item in the argument list like a for
  each.
* `watch ./8ball` will just run the program repeatedly instead of using
  a while loop. The default is 2 seconds. The -n argument will change
  the seconds
* Alternative terms for associative arrays:
  * dictionaries
  * maps
  * hash maps
  * hash tables
  * hashes
  * key/value pair
* Why would looking up a key/value pair not be safe for concurrency (in
  go)?
  * The algorithm that is used doesn't promise to keep track how often
    you are going to read that value. In other words `memoization`.
* Associative arrays in bash are not compound.
  * You cannot have an element be another array in bash. Other languages
    will allow for this.

