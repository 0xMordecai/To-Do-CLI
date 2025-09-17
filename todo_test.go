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

// TestComplete tests the Complete method of the List type
func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed.")
	}

	if !l[0].Done {
		t.Errorf("New task should be completed.")
	}
}

// TestDelete tests the Delete method of the List type
func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task ",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}

}
