# Boost Week 15 Notes
## Screen
* Allows for windows management in the terminal like tmux.
* Screen has been around for a long time. `sudo apt install -y screen`
  to install.
* To use screen commands, there is a leader key combination that must be
  used first `ctl+a` followed by the command.
  * `ctl+a+c` will create a new screen window.
  * `ctl+a+a` will cycle through windows.
  * `ctl+a+w` will display the windows active.
  * `ctl+a+[` will start copy mode. Use `vi` navigation to move cursor
    around and space bar to start to highlight section to copy and enter
    to copy highlighted section to buffer.
  * `ctl+a+]` will paste out the text that was copied from tmux.  
* `screen -r` will reattach to the last session
  * `screen -ls` will list out the screen sessions
* Default tmux config is different to screen but they are both
  multiplexers
* Tmux allows for split panes and additional configuration.
* `ctl+a+,` to rename window.
* `ctl+a+:` to manually execute a tmux command.
* `ctl+a+ctl+hold vi nav` to resize window pane.
* `ctl+a+r` to reload tmux config
* `ctl+a+t` to have a time screen
* Caution for running nested tmux sessions. There is a custom script
  created called `tmuxin` that would create a new nested tmux session.
* A reason for wanting nested tmux sessions is if you want a persistent
  pane to be present by using a separate configuration file. An example
  of this would be for streaming. You can have a persistent pane for the
  chat, date and time ect.
* The requirements for a basic configuration setup for any system are:
  .bashrc, .vimrc and .tmux.conf for most things we need.
