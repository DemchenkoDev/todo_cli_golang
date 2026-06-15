package scanner

import (
	"fmt"
	"toDoList/todo"

	"github.com/google/uuid"
)

// printPrompt prints the CLI prompt symbol.
func printPrompt() {
    fmt.Print("> ")
}

// printResult prints an error message returned by a command.
func printResult(result string) {
    fmt.Println("Error:", result)
}

// printExit prints a message before exiting the program.
func printExit() {
    fmt.Println("Exiting...")
}

// printHelp prints the list of available CLI commands.
func printHelp() {
    fmt.Println("Available commands:")
    fmt.Println("  add <title> <text>   - Add a new task")
    fmt.Println("  list                 - Show all tasks")
    fmt.Println("  done <UUID>          - Mark a task as done")
    fmt.Println("  del <UUID>           - Delete a task")
    fmt.Println("  events               - Show command history")
    fmt.Println("  help                 - Show this help message")
    fmt.Println("  exit                 - Exit the program")
}

// printAdd prints information about a newly created task.
// Accepts both ID and title because scanner.go passes two arguments.
func printAdd(id string, title string) {
    fmt.Println("Task added:")
    fmt.Println("  ID:    ", id)
    fmt.Println("  Title: ", title)
}

// printDone prints confirmation that a task was marked as completed.
func printDone(id string) {
    fmt.Println("Task marked as done:", id)
}

// printDel prints confirmation that a task was deleted.
func printDel(id string) {
    fmt.Println("Task deleted:", id)
}

// printTasks prints all tasks in the list with full details.
func printTasks(tasks map[uuid.UUID]todo.Task) {
    if len(tasks) == 0 {
        fmt.Println("No tasks found.")
        return
    }

    fmt.Println("Tasks:")
    fmt.Println("--------------------------------------------------")

    for id, t := range tasks {
        status := "Not done"
        if t.IsDone {
            status = "Done"
        }

        fmt.Println("ID:       ", id.String())
        fmt.Println("Title:    ", t.Title)
        fmt.Println("Text:     ", t.Text)
        fmt.Println("Created:  ", t.CreatedAt.Format("2006-01-02 15:04:05"))

        if t.DoneAt != nil {
            fmt.Println("Done at:  ", t.DoneAt.Format("2006-01-02 15:04:05"))
        } else {
            fmt.Println("Done at:   -")
        }

        fmt.Println("Status:   ", status)
        fmt.Println("--------------------------------------------------")
    }
}

// printEvents prints the command history.
func printEvents(events []Event) {
    if len(events) == 0 {
        fmt.Println("No events recorded.")
        return
    }

    fmt.Println("Events:")
    for _, e := range events {
        fmt.Printf("[%s] %s\n", e.Timestamp.Format("15:04:05"), e.Input)
    }
}
