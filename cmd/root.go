package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
// It's the 'tasker' command itself.
var rootCmd = &cobra.Command{
	Use:   "tasker",
	Short: "Tasker is a CLI for managing your to-do list.",
	Long: `Tasker is a command-line interface (CLI) tool to help you manage your to-do list.

You can add new tasks, list all existing tasks, and mark tasks as complete
right from your terminal. Your tasks are saved to a file in your home directory,
so they persist between sessions.

Examples:
  tasker add "Write the Go CLI tutorial"
  tasker list
  tasker complete 1`,
	// The root command itself doesn't do anything, so we don't need a 'Run' function.
	// If you type just 'tasker', Cobra will automatically display the help text.
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// The 'execute' function is the entry point of our command package.
	// It executes the root command. Cobra then figures out which sub-command
	// was called (e.g., 'add', 'list') and runs its corresponding 'Run' function.
	if err := rootCmd.Execute(); err != nil {
		// Cobra prints its own errors, but we exit with a non-zero status code
		// to indicate that an error occurred. This is standard practice for CLIs.
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

// The init function is a special Go function that runs before main().
// Here, we can set up flags and other configurations.
func init() {
	// In this simple application, we don't have any global flags that apply
	// to the root command itself or all sub-commands.
	//
	// If we wanted to add a persistent flag, like specifying a different
	// data file, we would do it here. For example:
	//
	// var dataFile string
	// rootCmd.PersistentFlags().StringVar(&dataFile, "file", "", "A different data file to use (default is $HOME/.tasks.json)")
	//
	// This flag would then be available for all commands, e.g., 'tasker list --file another-tasks.json'
}
