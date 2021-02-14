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


### Ask

```sh
$ eddie ask [behavior, symptom] -{d,w,m}

$ eddie ask buddha // returns a motivational message from Buddha
```

### Track

Eddie tracks your sleep, mood, weight, and much more!

```sh
$ eddie track sleep \
	-t 10:20 \
	-d 7 \
	-q 9 \
	-t one,two,three,four,five,six,seven \
	-l couch \
	-i

$ eddie track sleep \
	--time=10:20 \
	--duration=9 \
	--quality=8 \
	--tags=one,two,three,four,five \
	--location=couch \
	--important

$ eddie track weight --total=150

$ eddie track mood \
    -q 9 \
    --tags=happy,sad,one,three,five,seven,nine \
    -i

$ eddie track mood \
    -quality 9 \
    --tags=happy,sad,one,three,five,seven,nine \
    -important
    
$  eddie track seizure \
	--time="10:20" \
	--tags="one,two,three,four,five" \
	--location="strip club" \
	--important

$  eddie track seizure \
	-t 10:20 \
	--tags="one,two,three,four,five" \
	-l="strip club" \
	-i
```


Copyright
---------

Copyright (c) 2021, Rye Miller