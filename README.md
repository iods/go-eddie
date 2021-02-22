Eddie CLI
=========

A CLI for tracking routines and behaviors, with a UI for rendering your results in real time.


Description
-----------

I am in the command line a lot. I like data. I love Go right now. It feels right.


```
            __     __  __
 .-----..--|  |.--|  ||__|.-----.
 |  -__||  _  ||  _  ||  ||  -__|
 |_____||_____||_____||__||_____|

 Like velcro.

Usage:
  eddie [command]

Available Commands:
  ask         Ask eddie to remind you about tracking your routines daily, weekly, or monthly.
  help        Help about any command
  report      Report on your behaviors in the command line or in a dashboard.
  track       Track various routines and behaviors with some tags for filtering later.

Flags:
  -h, --help   help for eddie

Use "eddie [command] --help" for more information about a command.
```
eddie track mood 7 -tags="this,that,then,bag,of,chips" -stress=6 -quality=3 -emoji="joyful,happy,grateful,sad,loving,anxious,bored,frustrated,stressed)
https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

### Ask

```sh
$ eddie ask [behavior, symptom] -{d,w,m}

$ eddie ask buddha // returns a motivational message from Buddha
```

### Track

Eddie tracks your sleep, mood, weight, and much more!


## Development

### TermUI

### GORM, GoFakeit (to fill db)


Copyright
---------

Copyright (c) 2021, Rye Miller