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

func main() {

	//check show the manual if there is no arguments provided
	if len(os.Args) < 2 {
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
		return
	}

	tasks, err := LoadTasks()

	if err != nil {
		fmt.Printf("error loading tasks :%v", err)
	}

	switch os.Args[1] {
	case "list":
		{
			fmt.Println("All tasks:")

			for _, task := range tasks {
				fmt.Printf("%v-%v (%v) \n", task.ID, task.Description, task.Status)
			}
		}
	case "add":
		{
			if len(os.Args) < 3 {
				fmt.Println("Error: Please provide a task description")
				return
			}

			err := AddTask(&tasks, os.Args[2])

			if err != nil {
				fmt.Printf("error saving tasks :%v", err)
				return
			}

			fmt.Println("Task Saved!")
		}
	case "update":
		{
			if len(os.Args) < 4 {
				fmt.Println("Error: Please provide the task id and description")
				return
			}

			err := UpdateTask(&tasks, os.Args[2], os.Args[3])

			if err != nil {
				fmt.Printf("error saving tasks :%v", err)
				return
			}

			fmt.Println("Task Updated!")

		}
	case "delete":
		{
			if len(os.Args) < 3 {
				fmt.Println("Error: Please provide the task id ")
				return
			}

			err := DeleteTask(&tasks, os.Args[2])

			if err != nil {
				fmt.Printf("error Deleting task :%v", err)
				return
			}

			fmt.Println("Task deleted!")

		}
	case "version":
		{
			fmt.Println(version)
		}
	default:
		{
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
				task-tracker delete 1 done
				task-tracker list
				task-tracker list todo
		`)
		}
	}

}
