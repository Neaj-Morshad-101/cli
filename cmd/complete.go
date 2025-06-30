// file: cmd/complete.go
package cmd

import (
	"fmt"
	"strconv"

	"github.com/neaj-morshad-101/tasker-cli/db"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task ID]",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: You must provide the ID of the task to complete.")
			return
		}

		// Convert argument from string to integer
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Invalid task ID. Please provide a number.")
			return
		}

		tasks, err := db.ReadTasks()
		if err != nil {
			fmt.Println("Error reading tasks:", err)
			return
		}

		// Find the task with the given ID
		var taskFound bool
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Completed = true
				if err := db.WriteTasks(tasks); err != nil {
					fmt.Println("Error writing tasks:", err)
					return
				}
				fmt.Printf("Completed task: \"%s\"\n", tasks[i].Text)
				taskFound = true
				break
			}
		}

		if !taskFound {
			fmt.Printf("Error: Task with ID %d not found.\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
