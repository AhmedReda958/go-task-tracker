package main

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	StatusToDo TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone TaskStatus = "done"
)

type Task struct {
	ID int `json:"id"`
	Description string 	`json:"description"`
	Status TaskStatus `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func main() {
	newTask := Task{
		ID:1,
		Description: "Finsh task manager project",
		Status: StatusToDo,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	fmt.Println(newTask);

	newTask.Status = StatusDone

	fmt.Println(newTask);

}