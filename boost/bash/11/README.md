# Boost week 11 notes

## Review challenge:
* Create a new script from scratch. The script should print "Hello there"
* The logic must be put into a function call.

## Recap:
I used echo rather than printf. I should remember that I can pass
through the parameter when calling the function using "\$1"  which can be
used for parameter expansion as well.

## Notes:
* entr bash -c "clear;./greet Rob" <<< greet - This command will hot
  reload the script if there are any changes to it."
If bash syntax
```
if [[ condition ]]; then
  execution
fi
```

* What is the purpose of bash?
  * Bash is for rapid applications development and for testing ideas.
  * Write a script quickly that is most likely not going to stay around
    for a long time.
  * If it ever does become important, it should be migrated to a tool
    like go. Googles own bash style guide recommends to move to a
    different language if the bash script is 100 lines of code.

* Expansion
  * Substitues the expressions with something else.
  * \${1..10} will substitute for 1 to 10
  * \${variable}.md will substitute with the defined variable
  * Be careful of not using double quotes for variables that have spaces
    because bash will execute as separate commands space separated.
  * Tilde expansion will evaluate to \$HOME path for user.
  * \${#parameter} - parameter length used to see how many arguments you
    have.
  * \${parameter#* } used to exclude until the first instance of the
    third argument after '\*'. \${parameter## \*} The double hash will
    get the exclude the last instance. The inverse if true for % instead
    of #. Examples included in greet. This can be used over sed or awk
    because it is native bash.
  * \${parameter/search/replace} This should replace a term for the
    first instance.
* Prefer using () over \` \` because you cannot nest with \` \`
* Arithmetic Expansion - To do math in bash, you have to either use
    double parentheses or using bc. bc is not bash. This type of
    expansion does not need to \$ notation at all inside to reference
    variables or outside that are not expressions.
```
n=$((n++))
# Output 0
n=$((n++))
# Output 0
# The $ is not required
((n++))
# Output 1
# $ is required for expressions
local some=34
local other=20
echo "total: $((some + other))"
#Output: total: 54
```
* Bash does not do floating point math but bc can. \$((1/3)) will
    evaluate to 0.
* To evaluate this, you should use bc. This is the only safe way to do
  floating point math on any unix system.
```
#scale is the precision
$(bc <<< "scale=2;1/3")
#Ouput: .33
```
* bc can also be used to convert base numbers using 'obase' and 'ibase'.
  For example, using a ibase=8 and obase=2, you can use it visualise rwx
  for file permissions. Input of 644 will be 110100100 in binary.
* In bash, there is an environment variable called RANDOM which selects
  a random number between 0 and 32000.
```
# Select random number between 0 and 8.
echo $((RANDOM % 8))
# Select random number between 1 and 20.
echo $((RANDOM%20 + 1))
```
* Process substitution - takes the output of file descriptor using <.
  <<< redirects to stdin and <(...) gives you a file name. An example of
  this being used is the entr command.
```
entr bash -c "clear; ./greet" <<< "$(find . -name 'gr*')"
entr bash -c "clear; ./greet" < <(find . -name 'gr*')
```
* "\$@" pass through all of the arguments separated by a space
* "\$\*" concatenates  all of the arguments as one string
* Short circuit logic - only works with true. Each condition that is
  true is executed in order. The first condition that is false, stops
  the evaluation of the condition and the return is false.
* If the first condition is true, it will not read the next condition.
* Starting a script with `set -e` will stop the script when any error
  has been encountered. The default behaviour of shell is to carry on
  execution even after an error has been encountered.
  * Advised for `set -e` to be in script.
* `set -x` is used for debug purposes. This is print out the line that
  has been executed so it should only be used for development.
* `=~` to match against a regular expression. This should be done within
  double brackets. The regular expressions are extended type (basic,
  extended and pearl being the others).



