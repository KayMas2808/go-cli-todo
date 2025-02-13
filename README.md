# go-cli-todo

A simple command-line to-do list application written in Go.  This project is a great starting point for learning Go backend development concepts like working with data structures, file I/O, and basic command-line interaction.

## Features

*   Add new tasks
*   List existing tasks
*   Mark tasks as complete
*   Saves and loads tasks from a JSON file (`todos.json`)

## Getting Started

1.  **Prerequisites:** Make sure you have Go installed on your system. You can download it from [https://go.dev/](https://go.dev/).

2.  **Clone the Repository (Optional):** If you'd like to contribute or modify this project, you can clone the repository:

    ```bash
    git clone [https://github.com/yourusername/go-cli-todo](https://www.google.com/search?q=https://github.com/yourusername/go-cli-todo) // Replace with your repo URL
    ```

3.  **Navigate to the Directory:**

    ```bash
    cd go-cli-todo
    ```

4.  **Build the Application:**

    ```bash
    go build
    ```

## Usage

The `go-cli-todo` executable will be created in the project directory. You can then run it with the following commands:

```bash
./go-cli-todo <command> [arguments]
Available Commands:

add -task "Your task description": Adds a new task.  Replace "Your task description" with the actual task.

Bash

./go-cli-todo add -task "Buy groceries"
list: Lists all tasks with their status (completed or not).

Bash

./go-cli-todo list
complete -index <index>: Marks a task as complete. Replace <index> with the task number shown in the list command.

Bash

./go-cli-todo complete -index 1  // Completes the first task in the list
Example Workflow:

Bash

./go-cli-todo add -task "Learn Go"
./go-cli-todo add -task "Build a CLI app"
./go-cli-todo list
./go-cli-todo complete -index 2
./go-cli-todo list
Further Learning
This project is inspired by the concepts outlined in this roadmap: https://roadmap.sh/projects/task-tracker
