package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/0xMordecai/To-Do-CLI"
)

// Hardcoding the file name
const todoFileName = ".todo.json"

func main() {
	// Parsing command line Flags
	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()
	// Define an items list
	l := &todo.List{}

	// Use the Get() method to read to-do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// todo : V1
	// // Decide what to-do based on the number of arguments provided
	// switch {
	// // For no extra arguments, print the list
	// case len(os.Args) == 1:
	// 	// List current to do items
	// 	for _, item := range *l {
	// 		fmt.Println(item.Task) // --> to print only Task
	// 		//	fmt.Println(item)  // --> to print ALL info :=> fixed BUG
	// 	}
	// // Concatenate all provided arguments with a space and
	// // add to the list as an item
	// default:
	// 	// Concatenate all arguments with a space
	// 	item := strings.Join(os.Args[1:], " ")

	// 	// Add the task
	// 	l.Add(item)

	// 	// Save the new list
	// 	if err := l.Save(todoFileName); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// 	}
	// }

	// Decide what to do based on the provided flags
	switch {
	case *list:
		// list current to do items
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}

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

	}

}
