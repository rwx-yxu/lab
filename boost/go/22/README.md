#Boost notes week 22
* Logic errors can go undetected very easily. More so in interpreted
  languages than compiled because the compiler will catch logic errors
  during compile.

```
//Will not compile in Go because Go expects both sides of the statement
to be boolean
if name =="Lancelot" || "Galahad"

//Will execute in python
if name=="Lancelot" or "Galahad"
```
* Colours make testing difficult because std out will print their escape
  sequences. As well as if you want to pipe output to a file.
  * One way to deal with colours is to turn them off:
    * By turning each colour into a function call to detect the terminal
      environment before returning the colour sequence. This is much
      slower than a constant or var.
    * Change the colour sequences to var than constant to allow the
      colour variable to change to non.
* Go defaults to using the dependency listed in go.sum. Unless specified
  to use go.work instead.
* Go package dependencies are stored in `$GOPATH/pkg/mod/`. This
  directory stores the code of the dependencies listed  by their
  checksums.
* This is safe for go to default to the sourced version at the time of
  go get rather than automatically pull down the latest version from
  github. Node used to do this for the first couple years...
* How to tell if the program is being ran by a user or being ran by a
  program like in a test?
* Check and see if standard output (file descriptor 1 is standard out) is to a terminal for a user.
  * If it is not, then it is being fed into another program.
* Should only include terminal colours if you are sure that you are
  running an interactive terminal.
