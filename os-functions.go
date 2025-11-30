package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveTasks(tasks []Task) error {

	data, err := json.MarshalIndent(tasks, "", " ")

	if err != nil {
 		return fmt.Errorf("SaveTasks error: failed to marshal tasks to JSON: %v", err)
	}

	err = os.WriteFile(taskFile, data, 0644)

	if err != nil {
 		return fmt.Errorf("SaveTasks error: failed to write to file '%s': %v", taskFile, err)
	}
	return nil

}

func LoadTasks() ([]Task, error) {

	data, err := os.ReadFile(taskFile)

	if os.IsNotExist(err) {
		return []Task{}, nil
	}

	if err != nil {
 		return nil, fmt.Errorf("LoadTasks error: failed to read file '%s': %v", taskFile, err)
	}

	var tasks []Task

	if err := json.Unmarshal(data, &tasks); err != nil {
 		return nil, fmt.Errorf("LoadTasks error: failed to unmarshal JSON: %v", err)

	}

	return tasks, nil
}
