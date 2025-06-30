// file: cmd/list.go
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/neaj-morshad-101/tasker-cli/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadTasks()
		if err != nil {
			fmt.Println("Error reading tasks:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks! Why not add one with 'tasker add'?")
			return
		}

		fmt.Println("You have the following tasks:")

		// Using a tabwriter for nicely formatted columns
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "ID\tTASK\tSTATUS")
		fmt.Fprintln(w, "--\t----\t------")
		for _, task := range tasks {
			status := "Pending"
			if task.Completed {
				status = "✓ Completed"
			}
			fmt.Fprintf(w, "%d\t%s\t%s\n", task.ID, task.Text, status)
		}
		// Flush the writer to ensure all output is written to stdout
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
