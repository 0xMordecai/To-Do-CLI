package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/0xMordecai/To-Do-CLI"
)

// Hardcoding the file name
var todoFileName = ".todo.json"

func main() {
	// change flag.Usage() to display a custom message.==> Now that the user can get proper usage information.
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed for Organize ToDo Tasks\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2025@LCA DevOps Systems\n\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	// Parsing command line Flags
	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()
	// Check if the user defined the ENV VAR for costume nfile name
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}
	// Define an items list
	l := &todo.List{}

	// Use the Get() method to read to-do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the provided flags
	switch {
	case *list:
		// list current to do items
		fmt.Print(l)

	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *task != "":
		// Aff the task
		l.Add(*task)

		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		// Invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)

	}

}

// The getTask() function accepts as input the parameter r of type io.Reader interface
// and the parameter args, which consists of zero or more values of type string,
// represented by the ... operator preceding the parameter type
// The function getTask() returns a string and a potential error

func getTask(r io.Reader, args ...string) (string, error) {
	//verifies if any arguments were provided as the parameter args
	if len(args) > 0 {
		return strings.Join(args, ""), nil
	}
	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", nil
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}
	return s.Text(), nil
}
