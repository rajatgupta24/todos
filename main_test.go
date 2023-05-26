package main

import "testing"

// create a func that create a todo
// create a func that delete a todo
// Create a func to get todos

func TestAddTodos(t *testing.T) {
	todos := []string{"task 1"}
	want := []string{"task 1", "task 2"}

	got := AddTodos(todos, "task 2")

	for k := range want {
		if want[k] != got[k] {
			t.Errorf(`wanted:"%q" and got:"%q"`, want[k], got[k])
		}
	}
}
