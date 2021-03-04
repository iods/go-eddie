Eddie CLI
=========

A CLI for tracking routines and behaviors, with a cli dashboard for rendering your results in real time.


Description
-----------

I am in the command line a lot. I like data. I love Go right now. It feels right.


```

  _______ ______  ______  _____ _______
  |______ |     \ |     \   |   |______
  |______ |_____/ |_____/ __|__ |______

     the cli service dog 🐕 v0.1.0

Usage:
  eddie [command]

Available Commands:
  ask         Reminders for tracking activity
  help        Help about any command
  report      Generate reports for an activity
  track       Record and track patterns in sleep, mood, seizures, and weight.

Flags:
  -h, --help   help for eddie

Use "eddie [command] --help" for more information about a command.

```

### Getting Started

Simply clone the repository, set up an alias for eddie like `alias eddie=$GOBIN/eddie`, and then run:
```sh
$ go install && eddie install
```

Development
-----------

### Weight

You have only one main option with this record, and that is the average value (in pounds) of your weight,
taken over a 24 hr period. You are also given the `--important` flag for tagging it with significance.

You can track your weight using the following patterns:
```sh
$ eddie track weight 195
$ eddie track weight 195 -i, --important
```

You will see the record created return its ID, but nothing more.

You can view progress and trends with:
```sh
$ eddie report weight
```



#### Development

#### TermUI

#### GORM, GoFakeit (to fill db)


Copyright
---------

Copyright (c) 2021, Rye Miller