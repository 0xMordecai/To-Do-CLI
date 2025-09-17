package todo_test

import (
	"testing"

	todo "github.com/0xMordecai/To-Do-CLI.git"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "NAw Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}

}
