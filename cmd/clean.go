// file: cmd/clean.go
package cmd

import (
	"fmt"

	"github.com/neaj-morshad-101/tasker-cli/db"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Removes all completed tasks from your to-do list",
	Long: `The clean command finds all tasks that have been marked as complete
and permanently removes them from your task list. This is useful for keeping
your list tidy.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadTasks()
		if err != nil {
			fmt.Println("Error reading tasks:", err)
			return
		}

		// Create a new slice to hold only the tasks we want to keep (pending tasks)
		var tasksToKeep []db.Task

		for _, task := range tasks {
			if !task.Completed {
				tasksToKeep = append(tasksToKeep, task)
			}
		}

		// Calculate how many tasks were removed for the user message
		tasksRemoved := len(tasks) - len(tasksToKeep)

		if tasksRemoved == 0 {
			fmt.Println("No completed tasks to clean.")
			return
		}

		// Re-index the IDs of the remaining tasks to be sequential
		for i := range tasksToKeep {
			tasksToKeep[i].ID = i + 1
		}

		// Write the new, smaller slice of tasks back to the file
		if err := db.WriteTasks(tasksToKeep); err != nil {
			fmt.Println("Error writing tasks:", err)
			return
		}

		// Provide a helpful confirmation message
		// Handle pluralization for the message
		taskWord := "task"
		if tasksRemoved > 1 {
			taskWord = "tasks"
		}
		fmt.Printf("Successfully cleaned %d completed %s.\n", tasksRemoved, taskWord)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
