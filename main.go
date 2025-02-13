package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Todo struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func addTask(task string) {
	newTodo := Todo{Task: task, Completed: false}
	todos = append(todos, newTodo)
	saveTodos()
	fmt.Println("Task added: ", task)
}

func listTasks() {
	if len(todos) == 0 {
		fmt.Println("No tasks added.")
		return
	}
	for i, todo := range todos {
		status := "[ ]"
		if todo.Completed {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, todo.Task)
	}
}

func completeTask(ind int) {
	if ind < 1 || ind > len(todos) {
		fmt.Println("Invalid index")
		return
	}
	todos[ind-1].Completed = true
	saveTodos()
	fmt.Printf("Task %d completed: %s\n", ind, todos[ind-1].Task)
}

const todosFile = "todos.json"

func saveTodos() {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		fmt.Println("Error saving todos: ", err)
		return
	}
	err = os.WriteFile(todosFile, data, 0644)
	if err != nil {
		fmt.Println("Error saving todos: ", err)
	}
}

func loadTodos() {
	data, err := os.ReadFile(todosFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error loading todos: ", err)
		return
	}
	err = json.Unmarshal(data, &todos)
	if err != nil {
		fmt.Println("Error loading todos: ", err)
	}
}

func main() {
	loadTodos()

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addText := addCmd.String("task", "", "Task to add")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	completeIndex := completeCmd.Int("index", 0, "Index of task to complete")

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [arguments]")
		fmt.Println("Commands: add, list, complete")
		return
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addText == "" {
			fmt.Println("Task cannot be empty")
		} else {
			addTask(*addText)
		}
	case "list":
		listCmd.Parse(os.Args[2:])
		listTasks()
	case "complete":
		completeCmd.Parse(os.Args[2:])
		if *completeIndex == 0 {
			fmt.Println("Index is required.")
		} else {
			completeTask(*completeIndex)
		}
	default:
		fmt.Println("Invalid command: ", os.Args[1])
	}

}
