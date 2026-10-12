# Task Tracker CLI

A lightweight, command-line task management tool built in Go. This project is a completely dependency-free CLI application that stores task data locally in a JSON file. 

Inspired by the [roadmap.sh Task Tracker project](https://roadmap.sh/projects/task-tracker).

## Features

*   **Add Tasks:** Quickly add new tasks to your list.
*   **List Tasks:** View all tasks in a neatly formatted table.
*   **Filter by Status:** View tasks filtered by their current status (e.g., `done`, `todo`).
*   **Mark as Done:** Update the status of existing tasks.
*   **Local Storage:** All tasks are persisted locally in a `tasks.json` file.

## Prerequisites

*   [Go](https://go.dev/dl/) installed on your machine (version 1.16 or higher recommended).

## Installation & Setup

1. **Clone or download the repository:**
   ```bash
   git clone <your-repo-url>
   cd task-tracker
   ```

2. **Initialize the Go module (if not already done):**
   ```bash
   go mod init task-tracker
   ```

3. **Build the executable:**
   
   *On Windows (e.g., in VS Code terminal):*
   ```bash
   go build -o task-cli.exe main.go
   ```
   
   *On Linux/macOS:*
   ```bash
   go build -o task-cli main.go
   ```

## Usage

Once built, you can run the executable directly from your terminal. 

*Note: If you are using Windows PowerShell or Command Prompt, use `.\task-cli.exe`. On Linux/macOS, use `./task-cli`.*

### Add a Task
```bash
.\task-cli.exe add "Buy groceries"
.\task-cli.exe add "Finish Go project"
```

### List All Tasks
```bash
.\task-cli.exe list
```

### List Tasks by Status
```bash
.\task-cli.exe list done
.\task-cli.exe list todo
```

### Mark a Task as Done
Use the ID of the task you want to update (found via the `list` command):
```bash
.\task-cli.exe mark-done 1
```

## Data Storage

The application automatically creates and manages a `tasks.json` file in the same directory as the executable. You do not need to create this file manually.

## Next Steps / TODOs

*   [ ] Add the `update` command to change a task's description.
*   [ ] Add the `delete` command to remove a task completely.
*   [ ] Add a `mark-in-progress` command.