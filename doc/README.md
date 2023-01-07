Eddie - Documentation
=====================

Rationale
---------

Building an application that I can actually use. And documenting it.


Description
-----------
A tool written in Go for keeping track of all your projects and teammates deliverables.

A mix between a bulletin board and notebook. `go-standups` offers the following functionality:

* [ ] create, read, update, delete the core notes for a team
* [ ] generate docs based on env templates - encrypted creds
* [ ] set a reminder for daily, weekly, monthly updates
* [ ] export all notes and updates to one local file depending on the tag
    * i.e `drkctl add --client` goes to client.md
* [ ] randomly asks for a client to get an update from
    * alphabetical
    * JIRA statuses
    * as-provided
* [ ] each week, auto-commits and auto-merges, and pushes your notes to repo
* [ ] tags
* [ ] jira integration
* [ ] google calendar integration

---

- `drkctl add -n`: Add a new note.
- `d`: Delete a note. Author only.
- `e`: Edit a note. Author only.

### Getting Started

Simply clone the repository, set up an alias for standup like `alias standup=$GOBIN/standup`, and then run:
```sh
$ go install && standup install
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
