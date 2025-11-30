package main

import "fmt"

func handleList(tasks []Task, args []string) {
	var filteredTasks []Task
	var statusFilter TaskStatus
	var hasFilter bool

	if len(args) > 0 {
		parsedStatus, ok := parseTaskStatus(args[0])
		if !ok {
			fmt.Printf("Error: Invalid status '%s'. Valid statuses are: todo, in-progress, done\n", args[0])
			return
		}
		statusFilter = parsedStatus
		hasFilter = true
	}

	if hasFilter {
		for _, task := range tasks {
			if task.Status == statusFilter {
				filteredTasks = append(filteredTasks, task)
			}
		}
		fmt.Printf("Tasks with status '%s':\n", statusFilter)
	} else {
		filteredTasks = tasks
		fmt.Println("All tasks:")
	}

	if len(filteredTasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range filteredTasks {
		fmt.Printf("%v-%v (%v) \n", task.ID, task.Description, task.Status)
	}
}

func handleAdd(tasks *[]Task, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Please provide a task description")
		return
	}
	err := AddTask(tasks, args[0])
	if err != nil {
		fmt.Printf("main error: failed to save new task: %v\n", err)
		return
	}
	fmt.Println("Task Saved!")
}

func handleUpdate(tasks *[]Task, args []string) {
	if len(args) < 2 {
		fmt.Println("Error: Please provide the task id and description")
		return
	}
	err := UpdateTask(tasks, args[0], args[1])
	if err != nil {
		fmt.Printf("main error: failed to update task: %v\n", err)
		return
	}
	fmt.Println("Task Updated!")
}

func handleDelete(tasks *[]Task, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Please provide the task id")
		return
	}
	err := DeleteTask(tasks, args[0])
	if err != nil {
		fmt.Printf("main error: failed to delete task: %v\n", err)
		return
	}
	fmt.Println("Task deleted!")
}

func handleStatusUpdate(tasks *[]Task, args []string) {
	if len(args) < 2 {
		fmt.Println("Error: Please provide the task id and status")
		return
	}
	err := UpdateTaskSatus(tasks, args[0], args[1])
	if err != nil {
		fmt.Printf("main error: failed to update task: %v\n", err)
		return
	}
	fmt.Println("Task Updated!")
}
