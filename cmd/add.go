// file: cmd/add.go
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/neaj-morshad-101/tasker-cli/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Adds a new task to your to-do list",
	Run: func(cmd *cobra.Command, args []string) {
		// Join all arguments to form the task text
		taskText := strings.Join(args, " ")
		if taskText == "" {
			fmt.Println("Error: You must provide a description for the task.")
			return
		}

		// Read existing tasks
		tasks, err := db.ReadTasks()
		if err != nil {
			fmt.Println("Error reading tasks:", err)
			return
		}

		// Create a new task
		newTask := db.Task{
			ID:        len(tasks) + 1, // Simple ID assignment
			Text:      taskText,
			Completed: false,
			CreatedAt: time.Now(),
		}

		// Add the new task to the slice
		tasks = append(tasks, newTask)

		// Write the updated tasks list back to the file
		if err := db.WriteTasks(tasks); err != nil {
			fmt.Println("Error writing tasks:", err)
			return
		}

		fmt.Printf("Added new task: \"%s\"\n", taskText)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
