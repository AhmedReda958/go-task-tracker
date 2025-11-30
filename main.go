package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

	

	tasks, err := loadTasks()

	if err != nil {
		fmt.Printf("error loading tasks :%v", err)
	}

	switch os.Args[1] {
	case "list":
		{
			fmt.Println("All tasks:")

			for index, task := range tasks {
				fmt.Printf("%v-%v (%v) \n", index, task.Description, task.Status)
			}
		}
	case "add":
		{
			if len(os.Args) < 3 {
				fmt.Println("Error: Please provide a task description")
				return
			}

			content := os.Args[2]

			newTask := Task{
				ID:          len(tasks) + 1,
				Description: content,
				Status:      StatusToDo,
				CreatedAt:   time.Now().Local(),
				UpdatedAt:   time.Now().Local(),
			}

			tasks = append(tasks, newTask)

			err = saveTasks(tasks)

			if err != nil {
				fmt.Printf("error saving tasks :%v", err)
			}

			fmt.Println("Task Saved!")
		}
	case "update":
		{
			if len(os.Args) < 4 {
				fmt.Println("Error: Please provide the task id and description")
				return
			}

			idStr,description := os.Args[2],os.Args[3]

			id, err := strconv.Atoi(idStr)

			if err != nil {
				fmt.Println("Error: Invalid task ID")
				return
			}

			task := tasks[id]
			task.Description =description
			task.UpdatedAt= time.Now().Local()

			tasks[id]= task
			err = saveTasks(tasks)

			if err != nil {
				fmt.Printf("error saving tasks :%v", err)
			}

			fmt.Println("Task Updated!")

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

func saveTasks(tasks []Task) error {

	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
		return fmt.Errorf("error marsharling JSON: %v", err)
	}

	err = os.WriteFile(taskFile, data, 0644)

	if err != nil {
		return fmt.Errorf("error writing files: %v", err)
	}
	return nil

}

func loadTasks() ([]Task, error) {

	data, err := os.ReadFile(taskFile)

	if os.IsNotExist(err) {
		return []Task{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var tasks []Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)

	}

	return tasks, nil
}
