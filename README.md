[![Build Status](https://travis-ci.org/thash/asana.svg?branch=master)](https://travis-ci.org/thash/asana)

OverView
=========================================

[Asana](https://asana.com/) command line client implemented in Go.


Install
=========================================

Requirements: go

### Mac OS X

    $ brew tap thash/asana
    $ brew install asana


### Others

    $ go get github.com/thash/asana


Usage
=========================================

    $ asana help

    NAME:
       asana - asana cui client ( https://github.com/thash/asana )

    USAGE:
       asana [global options] command [command options] [arguments...]

    VERSION:
       x.x.x

    COMMANDS:
       config, c            Asana configuration. Your settings will be saved in ~/.asana.yml
       workspaces, w        get workspaces
       tasks, ts            get tasks
       task, t              get a task
       comment, cm          Post comment
       done                 Complete task
       due                  set due date
       browse, b            open a task in the web browser
       help, h              Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --help, -h           show help
       --version, -v        print the version


### Configure


    $ asana config
    visit: http://app.asana.com/-/account_api
      Settings > Apps > Manage Developer Apps > Personal Access Tokens
      + Create New Personal Access Token

    paste your Personal Access Token: _ <Copy Token from URL above and paste it.>

![](https://raw.githubusercontent.com/thash/asana/images/token.png)

When you paste valid token, your workspaces will be displayed.

    2 workspaces found.
    [0]    4444444444444 My Project
    [1]     999999999999 Work

    Choose one out of them: _

Select one workspace. Configurations are saved in `~/.asana.yml`.

    $ cat ~/.asana.yml

    personal_access_token: 0/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    workspace: 4444444444444


### Tasks

`asana tasks` or `asana ts` list your tasks.

    $ asana ts

    0 [ 2014-08-13 ] Write README.
    1 [ 2014-08-18 ] Buy gift for coworkers.
    2 [ 2014-08-29 ] Read "Unweaving the Rainbow".
    3 [            ] haircut

`asana task <index>` or `asana t <index>` shows the task in detail. When you run it without index, top of the tasks list will be used.

`-v` option adds comments and modification histories to the output.

    $ asana t -v 0

    [ 2014-08-13 ] Write README.
    --------
    Write README.md for Asana Cli project.

    ----------------------------------------

    assigned to you (2014-07-07T05:31:18.278Z)
    --------
    changed the name to "Write README." (2014-07-18T08:52:57.020Z)
    --------
    changed the due date to August 8 (2014-08-04T10:33:07.168Z)
    --------
    How about progress?
    by Lain Iwakura (2014-08-10T04:17:57.741Z)
    --------
    moved from Piyo to Hoge (2014-08-11T02:02:53.051Z)
    --------
    No progress.
    by Hash (2014-08-11T01:21:38.014Z)
    --------
    moved from Hoge to Piyo (2014-08-11T02:02:53.051Z)
    --------
    changed the due date to August 13 (2014-08-11T10:30:39.785Z)


### Complete, set due on a task

To complete task, use `asana complete <index>` or `asana done <index>`.

    $ asana done 12

To change(or newly set) due date, use `asana due <index> <due_date>`.

    $ asana due 5 2014-08-21

Or, `today` or `tomorrow`.

    $ asana due 5 today


### Comment

`asana comment <index>` or `asana cm <index>` enable you to post new comment for the task.

    $ asana cm 2

This command opens editor. Write comment, save and close.

![](https://raw.githubusercontent.com/thash/asana/images/cmt.png)

You can change editor by updating `$EDITOR` environment variable.


### Open a task in the browser

`asana browse <index>` or `asana b <index>` will open task in browser.

    $ asana browse 1
    // => open browser


TODO
=========================================

See [Issues](https://github.com/thash/asana/issues)
