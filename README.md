OverView
=========================================

[Asana](https://asana.com/) command line client implemented in Go.


Install
=========================================

Requirements: go

### Mac OS X

    $ brew tap memerelics/asana
    $ brew install asana


### Others

    $ go get github.com/memerelics/asana


Usage
=========================================

    $ asana help

    NAME:
       asana - asana cui client
    
    USAGE:
       asana [global options] command [command options] [arguments...]
    
    VERSION:
       0.0.1
    
    COMMANDS:
       config, c            Asana configuration. Your settings will be saved in ~/.asana.yml
       workspaces, w        get workspaces
       tasks, ts            get tasks
       task, t              get a task
       comment, cm          Post comment
       help, h              Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
       --help, -h           show help
       --version, -v        print the version


### Configure


    $ asana config
    visit: http://app.asana.com/-/account_api
    paste your api_key: _ <Copy API KEY from URL above and paste it.>

![](https://raw.githubusercontent.com/memerelics/asana/images/key.png)

When you paste valid api key, your workspaces will be displayed.

    2 workspaces found.
    [0]    4444444444444 My Project
    [1]     999999999999 Work
    
    Choose one out of them: _

Select one workspace.

Configurations are saved in `~/.asana.yml`.

    $ cat ~/.asana.yml
    
    api_key: 1xxxxxxx.xxxxxxxxxxxxxxxxxxxxxug
    workspace: 4444444444444


### Tasks

`asana tasks` or `asana ts` list your tasks.

    $ asana ts

    15384078744123 [ 2014-08-13 ] Write README.
    12869233163655 [ 2014-08-18 ] Buy gift for coworkers.
    14445736269211 [ 2014-08-29 ] Read "Unweaving the Rainbow".
    15232434010512 [            ] haircut

`asana task <task_id>` or `asana t <tasi_id>` shows the task in detail. When you run it without `task_id`, top of the tasks list will be used.

`-v` option adds comments and modification histories to the output.

    $ asana t -v 15384078744123

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


### Comment

`asana comment <task_id>` or `asana cm <task_id>` enable you to post new comment for the task.

    $ asana cm 15384078744123

This command opens editor. Write comment, save and close.

You can change editor by updating `$EDITOR` environment variable.


TODO
=========================================

* Run faster
* Select one task by index(`0, 1, 2...`) instead of long task_id (like `15384078744123`).
* Create new task
* Update(complete, set due date, edit title and notes etc) tasks
* Cancel Comment

