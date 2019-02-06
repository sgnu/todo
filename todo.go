package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	tasks := readFromFile()
	app := cli.NewApp()
	app.Name = "todo"
	app.Usage = "a todo list"
	app.Version = "1.0"

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a task to the list",
			Action: func(c *cli.Context) error {
				var buf bytes.Buffer //	Create a string builder to take the "args" and creates one long string
				buf.WriteString(c.Args().First())
				for i := 1; i < len(c.Args()); i++ {
					buf.WriteString(" " + c.Args().Get(i))
				}
				tasks = append(tasks, buf.String())
				sort.Strings(tasks)
				jsoninfo, _ := json.Marshal(tasks)
				writeToFile(jsoninfo)
				var arr []string
				_ = json.Unmarshal([]byte(jsoninfo), &arr)
				return nil
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "Complete a task on the list at an index or at the beginning",
			Action: func(c *cli.Context) error {
				sort.Strings(tasks)
				kb := bufio.NewReader(os.Stdin)
				index, _ := strconv.Atoi(c.Args().First())
				fmt.Print("You will be completing: " + tasks[index] + "\nare you sure? (y/n)")
				input, _ := kb.ReadByte()
				if input == 'y' {
					tasks = append(tasks[:index], tasks[index+1:]...)
					jsoninfo, _ := json.Marshal(tasks)
					writeToFile(jsoninfo)
					fmt.Println("Completed")
				} else {
					fmt.Println("Not completed")
				}
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List all the tasks on the list",
			Action: func(c *cli.Context) error {
				sort.Strings(tasks)
				for i, task := range tasks {
					fmt.Println(i, task)
				}
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Use the help to use this app")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func menu() {

}

//	Auxiliary functions

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFromUser() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter something: ")
	data, err := reader.ReadBytes('\n')
	check(err)
	return data
}

func readFromFile() []string {
	f, err := os.Open("/tmp/tasks")
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var arr []string

	_ = scanner.Scan()
	jsontext := scanner.Bytes()

	_ = json.Unmarshal(jsontext, &arr)
	return arr
}

func writeToFile(data []byte) {
	f, err := os.Create("/tmp/tasks")
	check(err)

	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.Write(data)

	writer.Flush()
}
