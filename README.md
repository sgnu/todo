# todo
CLI todo app written in go
## Installation
With go already installed, you can just run:

`go get github.com/sgnu/todo`

Make sure your `PATH` includes `$GOPATH/bin` to quickly run commands.
## Usage
There are 4 commands in todo:

1.	add, a
2.	complete, c
3.  edit, e
4.	list, l

#### Add
This command adds a new task to the list:
`todo add This is a new task` or `todo a This is a new task`

New tasks are added to the list and automatically sorted as strings

#### Complete
This command removes a task from the list. An index (starts at 0) can be specified, but otherwise removes the first task: `todo complete 1` or `todo c 1`

Arguments that are not integers are interpretted as 0

#### Edit
This command changes the name of a task on the list at a given index: `todo edit 1 new name for the task` or `todo e 1 new name for the task`

#### List
This command prints out all tasks on the list including their index: `todo list` or `todo l`

The list is automatically sorted before printing