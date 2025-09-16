package todo

import "time"

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
