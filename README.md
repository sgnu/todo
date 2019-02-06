# todo
CLI todo app written in go
## Installation
With go already installed, you can just run:

`go get github.com/sgnu/todo`

Make sure your `PATH` includes `$GOPATH/bin` to quickly run commands.
## Usage
There are 3 commands in todo:

1.	add, a
2.	complete, c
3.	list, l

##### Add
This command adds a new task to the list:
```todo add This is a new task``` or ```todo a This is a new task```

New tasks are added to the end of the list. Currently there is not an option to specify the index.
