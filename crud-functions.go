package main

import (
	"fmt"
	"strconv"
	"time"
)

func findTask(tasks []Task, id int) int {
	foundAt := -1

	for i := range tasks {

		if tasks[i].ID == id {
			foundAt = i
			break
		}
	}

	return foundAt
}

func getNextTaskID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(tasks *[]Task, content string) error {
	newTask := Task{
		ID:          getNextTaskID(*tasks),
		Description: content,
		Status:      StatusToDo,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   time.Now().Local(),
	}

	*tasks = append(*tasks, newTask)

	err := SaveTasks(*tasks)

	return err
}

func UpdateTask(tasks *[]Task, taskId string, description string) error {

	id, err := strconv.Atoi(taskId)

	if err != nil {
		return fmt.Errorf("UpdateTask error: invalid task id '%s': %v", taskId, err)

	}

	taskIndex := findTask(*tasks, id)
	if taskIndex == -1 {
		return fmt.Errorf("UpdateTask error: task with ID %d not found", id)

	}
	task := (*tasks)[taskIndex]
	task.Description = description
	task.UpdatedAt = time.Now().Local()

	(*tasks)[taskIndex] = task

	err = SaveTasks(*tasks)

	return err
}

func DeleteTask(tasks *[]Task, taskId string) error {
	id, err := strconv.Atoi(taskId)

	if err != nil {
		return fmt.Errorf("DeleteTask error: invalid task id '%s': %v", taskId, err)

	}

	taskIndex := findTask(*tasks, id)
	if taskIndex == -1 {
		return fmt.Errorf("DeleteTask error: task with ID %d not found", id)

	}

	*tasks = append((*tasks)[:taskIndex], (*tasks)[taskIndex+1:]...)

	err = SaveTasks(*tasks)

	return err
}
