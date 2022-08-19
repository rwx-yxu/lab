# Week 13 - bash configuration
* Use initial `.` to hide a file or directory.
* ls -la to see all hidden.
* Be careful of who you show your dot files to.
* `.profile` - Original login script. There is a difference between
  "login/interactive shell". Might be useful for mac.
* `.bashrc` - main bash configuration file
* `.config` - user (local) configurations (`os.UserConfigDir`)
* `.cache` - write things temporarily that are meant to go away. 
* `.local` - localized tree similar to /. Modernised place to put
  localised executables.
* `/etc/skel` is the exact copy of a new users profile. This contains
  the defaults.
* `cp -L` follow symbolic links
* `exec bash -l` - Reinitialize `.bashrc` file.
* Use \ to disable aliases and functions
* `kill process id` to kill a running process. `ps -ef |grep username`
  will find the running processes by the user name.
* `$-` returns the bash options such as if it is interactive.
* Use `uname` to see information about the system.
* `source` is the same as typing out every line of a file in that
  location as is. This is a source of security flaws.
* `set -o` to change the value of a shell option
* Bash only settings are set with `shopt`. Legacy settings are set from
  `set`
  * `$-` is the short form of `$BASHOPTS` which can be identified on
    the set help page
* Use `stty stop undef` to disable accidental terminal suspend
* `jobs` will show the running programs running in the back ground that
  ares suspended by ^Z.
* To use `set -o vi`, you will need to `esc` after the command and then
  you can use vi keys to search history.
* A bash function can only be executed from bash. This is why it may be
  more preferable to use a script instead so that it can be called from
  a separate program like go.

