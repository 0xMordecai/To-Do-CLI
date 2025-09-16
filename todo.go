package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// item struct represents a ToDo item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of ToDo items
type List []item

// Add creates a new todo item and appends it to the list
// --> (l *List) is a reciever that will attach Add() method to List type ;
// You can use any valid Go identifier,
// but it’s common to use the first letter of the type name as the identifier—in this case,l (lowercase L) for List
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	// 	Note that you need to dereference the pointer to the List type with *l in the
	// append call to access the underlying slice.
	*l = append(*l, t)
}

// Complete method marks a ToDo item as completed by
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes a ToDo item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	*l = append(ls[:i-1], ls[i:]...)

	// 	ls[:i-1]: Gets all items before the one to delete.
	//  ls[i:]: Gets all items after the one to delete.
	//  append(ls[:i-1], ls[i:]...): Concatenates the two slices, effectively removing the item at index i-1.
	// *l = ...: Updates the original list (since l is a pointer).

	return nil
}

// Save Method encodes the List as JSON and saves it
// using the provided file name

func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
}
