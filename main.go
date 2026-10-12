package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

const dataFile = "tasks.json"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "todo", "in-progress", "done"
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Printf("Please provide a task description!")
			return
		}
		fmt.Println(os.Args[0])
		addTask(os.Args[2])
	case "list":
		status := ""
		if len(os.Args) == 3 {
			status = os.Args[2]
		}
		listTasks(status)
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID.")
			return
		}
		updateStatus(os.Args[2], "done")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: task-cli <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  add <description>      Add a new task")
	fmt.Println("  list [status]          List all tasks, optionally filtered by status")
	fmt.Println("  mark-done <id>         Mark a task as done")
}

// --- File I/O Helpers ---

func loadTasks() ([]Task, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return []Task{}, nil // Return empty list if file doesn't exist
	}

	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

// --- Command Implementations ---

func addTask(description string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	// Generate new ID based on the highest existing ID
	newID := 1
	for _, t := range tasks {
		if t.ID >= newID {
			newID = t.ID + 1
		}
	}

	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving task:", err)
		return
	}
	fmt.Printf("Task added successfully (ID: %d)\n", newID)
}

func listTasks(filterStatus string) {
	tasks, err := loadTasks()
	if err != nil || len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	// Use tabwriter to create a neatly formatted table
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tDescription\tStatus\tCreated")
	fmt.Fprintln(w, "--\t-----------\t------\t-------")

	for _, t := range tasks {
		if filterStatus == "" || t.Status == filterStatus {
			dateStr := t.CreatedAt.Format("2006-01-02 15:04")
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", t.ID, t.Description, t.Status, dateStr)
		}
	}
	w.Flush()
}

func updateStatus(idStr, newStatus string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error: Task ID must be a number.")
		return
	}

	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Error: Task with ID %d not found.\n", id)
		return
	}

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	fmt.Printf("Task %d marked as %s.\n", id, newStatus)
}
