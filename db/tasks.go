// file: db/tasks.go
package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Task represents a single to-do item
type Task struct {
	ID        int
	Text      string
	Completed bool
	CreatedAt time.Time
}

// tasksFilePath returns the path to the tasks.json file.
func tasksFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get user home directory: %w", err)
	}
	return filepath.Join(home, ".tasks.json"), nil
}

// ReadTasks reads all tasks from the tasks.json file.
// If the file doesn't exist, it returns an empty slice.
func ReadTasks() ([]Task, error) {
	filePath, err := tasksFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		// If the file doesn't exist, it's not an error.
		// We'll just return an empty slice of tasks.
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("could not read tasks file: %w", err)
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("could not parse tasks file: %w", err)
	}
	return tasks, nil
}

// WriteTasks writes a slice of tasks to the tasks.json file.
func WriteTasks(tasks []Task) error {
	filePath, err := tasksFilePath()
	if err != nil {
		return err
	}

	// Marshal the data with indentation for readability
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal tasks: %w", err)
	}

	// Write the data to the file, creating it if necessary with 0644 permissions
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("could not write to tasks file: %w", err)
	}

	return nil
}
