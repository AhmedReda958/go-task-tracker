package main

import (
	"encoding/json"
	"fmt"
	"os"
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

const taskFile ="tasks.json"

func main() {
	tasks,err := loadTasks()

	if err !=nil {
		fmt.Printf("error loading tasks :%v", err)
	}

	fmt.Printf("alltasks: %v",tasks)

	newTask := Task{
		ID:1,
		Description: "Finsh task manager project",
		Status: StatusToDo,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	tasks = append(tasks, newTask)

	err=saveTasks(tasks)

	if err !=nil {
		fmt.Printf("error saving tasks :%v", err)
	}


	taskObject,_ := json.Marshal(newTask)
	fmt.Println(string(taskObject));

	newTask.Status = StatusDone

	fmt.Println(newTask);

}


func saveTasks (tasks []Task) error {

	data,err := json.MarshalIndent(tasks,""," ")
	
	if err != nil {
		return fmt.Errorf("error marsharling JSON: %v",err)
	}

	err = os.WriteFile(taskFile,data,0644);

	if err != nil {
		return fmt.Errorf("error writing files: %v",err)
	}
	return nil
	
}

func loadTasks () ([]Task , error) {

	data,err := os.ReadFile(taskFile)

	if os.IsNotExist(err) {
		return []Task{} , nil
	}

	if err != nil {
		return nil, fmt.Errorf("error reading file: %v",err)
	}

	var tasks []Task

	if err:=json.Unmarshal(data,&tasks); err!=nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v",err)
		
	}

	return tasks,nil
}