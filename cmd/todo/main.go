package main

import (
	"fmt"
	"os"

	todo "github.com/0xMordecai/To-Do-CLI"
)

// Hardcoding the file name
const todoFileName = ".todo.json"

func main() {
	// Define an items list
	l := &todo.List{}

	// Use the Get() method to read to-do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to-do based on the number of arguments provided
	switch {
	// For no extra arguments, print the list
	case len(os.Args) == 1:
		// List current to do items
		for _, item := range *l {
			fmt.Println(item)
		}

	}
}
