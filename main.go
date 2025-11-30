package main

import (
	"fmt"
	"os"
	"time"
)

type TaskStatus string

const (
	StatusToDo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

const taskFile = "tasks.json"
const version = "v1.0"

func printHelpManual() {
	fmt.Printf(`
			Usage: task-tracker <command> [arguments]

			Commands:
				add <description>    Add a new task
				update <id> <status> Update task status (todo, in-progress, done)
				delete <id>          Delete a task
				list [status]        List all tasks or filter by status

			Examples:
				task-tracker add "Buy groceries"
				task-tracker update 1 done
				task-tracker delete 1 
				task-tracker list
				task-tracker list todo
		`)
}

func main() {

	//check show the manual if there is no arguments provided
	if len(os.Args) < 2 {
		printHelpManual()
		return
	}

	tasks, err := LoadTasks()

	if err != nil {
		fmt.Printf("main error: failed to load tasks: %v\n", err)
	}

	switch os.Args[1] {
	case "list":
		handleList(tasks)
	case "add":
		handleAdd(&tasks, os.Args[2:])
	case "update":
		handleUpdate(&tasks, os.Args[2:])
	case "delete":
		handleDelete(&tasks, os.Args[2:])
	case "version":
		fmt.Println(version)
	default:
		printHelpManual()

	}

}
