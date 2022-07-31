# Notes

* Basic bash/shell data types: number/ineger, string and boolean
  * There are no floats in bash.
* Use declare to staticlly set a data type
  * declare -i will set an int
* Bash determines data types as 'duck typing'. It will iterpret the
  data type during executing. This is okay for writing quick
  applications but horrible for production environments.
* To check the type of a variable, you can use 'declare -p'
* Bash 4.0 and above support associative arrays but you have to
  declare them
* Parameter expansion starts with ${}. This can be used in place of sed
  and awk.
* '$#' is the number of arguements past through. Even a '' will count as
  one argument.
* You can also set default and assign default variables by this syntax
  ': "${PAGER=more}"'. This would save repeating the variable name
  for example 'declare PAGER=${PAGER:=more}'
* Parameter expansion also has it's own double quote scope as well so
  you don't have to escape them.

