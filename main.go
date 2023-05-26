package main

import "fmt"

func getTodos(todos []string) {
	for k, v := range todos {
		fmt.Println(k+1, v)
	}
	fmt.Println("*****")
}

func AddTodos(todos []string, todo string) []string {
	todos = append(todos, todo)
	return todos
}

func deleteTodos(todos []string, id int) []string {
	for k := range todos {
		if k+1 == id {
			return append(todos[:k], todos[k+1:]...)
		}
	}
	return todos
}

func main() {
	todos := []string{"Learn golang", "Test the app"}

	getTodos(todos)

	todos = AddTodos(todos, "Write a joke about go")
	getTodos(todos)

	todos = deleteTodos(todos, 1)
	getTodos(todos)
}
