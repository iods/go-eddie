Eddie CLI
=========

A CLI for tracking routines and behaviors, with a UI for rendering your results in real time.

Description
-----------

I am in the command line a lot. I like data. I love Go right now. It feels right.

Actors are just task runners.
Reports are generated in context, by their actors, keeping things isolated (recursion)
Tracking is a task, i.e. data collection, by an actor


```
eddie ask [behavior, symptom] -{d,w,m}

eddie ask buddha // returns a motivational message from Buddha, -lucky flag includes with lucky numbers
```


```
            __     __  __
 .-----..--|  |.--|  ||__|.-----.
 |  -__||  _  ||  _  ||  ||  -__|
 |_____||_____||_____||__||_____|

 A man's best friend.

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


Copyright
---------

Copyright (c) 2021, Rye Miller