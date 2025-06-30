# Tasker CLI

[![Go Version](https://img.shields.io/github/go-mod/go-version/neaj-morshad-101/tasker-cli?style=for-the-badge)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

> A simple but powerful command-line interface (CLI) for managing your to-do list right from the terminal.

Tasker is built with Go and Cobra to be fast, cross-platform, and easy to use. Never leave your terminal to manage your daily tasks again!

## Demo

Here is a quick demonstration of the core workflow:

![Tasker CLI Demo](./demo.gif)
*(To create a GIF like this, you can use tools like [asciinema](https://asciinema.org/) and `agg`)*

Alternatively, here's a text-based preview:

```bash
# Add a few new tasks
$ tasker add "Write the project README file"
Added new task: "Write the project README file"
$ tasker add "Deploy the new feature"
Added new task: "Deploy the new feature"

# List all tasks
$ tasker list
ID    TASK                           STATUS
--    ----                           ------
1     Write the project README file  Pending
2     Deploy the new feature         Pending

# Mark a task as complete by its ID
$ tasker complete 1
Completed task: "Write the project README file"

# List tasks again to see the change
$ tasker list
ID    TASK                           STATUS
--    ----                           ------
1     Write the project README file  ✓ Completed
2     Deploy the new feature         Pending

# Clean up by removing all completed tasks
$ tasker clean
Successfully cleaned 1 completed task.

# The list is now clean and re-indexed
$ tasker list
ID    TASK                           STATUS
--    ----                           ------
1     Deploy the new feature         Pending
```

## Features

- **Add:** Quickly add new tasks to your list.
- **List:** View all your tasks in a clean, formatted table.
- **Complete:** Mark tasks as complete using their ID.
- **Clean:** Remove all completed tasks to keep your list tidy.
- **Persistent:** Your tasks are saved to a `.tasks.json` file in your home directory, so they are always there when you come back.
- **Cross-Platform:** Works on macOS, Linux, and Windows.

## Installation

### Option 1: Using `go install` (For Go developers)

If you have Go installed and configured on your system, you can install `tasker-cli` with a single command.

```bash
go install github.com/neaj-morshad-101/tasker-cli@latest
```
This will download the source, compile it, and place the `tasker-cli` binary in your Go `bin` directory (`$HOME/go/bin` by default), which should be in your system's `PATH`.

### Option 2: From Pre-compiled Binaries (Recommended)

For most users, the easiest way is to download a pre-compiled binary for your operating system.

1.  Go to the [**Releases Page**](https://github.com/neaj-morshad-101/tasker-cli/releases) on GitHub.
2.  Download the appropriate archive for your system (e.g., `tasker-cli_linux_amd64.tar.gz` or `tasker-cli_windows_amd64.zip`).
3.  Extract the `tasker-cli` (or `tasker-cli.exe`) binary.
4.  Move the binary to a directory that is in your system's `PATH`. On Linux/macOS, `/usr/local/bin` is a great choice.
    ```bash
    # For Linux / macOS
    mv tasker-cli /usr/local/bin/
    ```

## Usage

Here are the available commands. You can also use `tasker --help` or `tasker [command] --help` to see more details.

### Main Commands

| Command | Description | Example |
| :--- | :--- | :--- |
| **`tasker add [task]`** | Adds a new task to your list. | `tasker add "Buy milk"` |
| **`tasker list`** | Lists all of your tasks. | `tasker list` |
| **`tasker complete [ID]`** | Marks a task as complete by its ID. | `tasker complete 2` |
| **`tasker clean`** | Removes all completed tasks. | `tasker clean` |

## Development

Interested in contributing? Here's how to get the project set up for development.

### Prerequisites

- [Go](https://go.dev/dl/) (version 1.21 or newer)
- [Git](https://git-scm.com/)

### Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/neaj-morshad-101/tasker-cli.git
    ```

2.  **Navigate to the project directory:**
    ```bash
    cd tasker-cli
    ```

3.  **Build the binary:**
    This command compiles the code and creates an executable named `tasker` in the current directory.
    ```bash
    go build -o tasker
    ```

4.  **Run for testing:**
    Execute the binary directly from the project directory.
    ```bash
    ./tasker list
    ```

## License

This project is distributed under the MIT License. See the `LICENSE` file for more information.