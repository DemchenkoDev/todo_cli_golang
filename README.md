# Todo CLI (v2)

A simple command-line Todo application written in Go.  
This version implements UUID-based task management, command parsing, event logging, and a clean CLI output.

This is an early working version (v2).  
Future versions will introduce better UX, internal architecture, command mapping, and more.

---

## Features

- Add tasks with title and text
- List all tasks with timestamps
- Mark tasks as done
- Delete tasks
- View command history
- UUID-based task identification
- Clean and structured CLI output

---

## Commands

add <title> <text> - Add a new task
list - Show all tasks
done <UUID> - Mark a task as done
del <UUID> - Delete a task
events - Show command history
help - Show help message
exit - Exit the program

---

## Example Usage

add BuyMilk 2 liters
Task added:
ID: 018f3b5e-9c3a-7c2b-8e3a-9f1c2d3b4e5f
Title: BuyMilk

list
Tasks:

ID: 018f3b5e-...
Title: BuyMilk
Text: 2 liters
Created: 2026-06-15 21:14:00
Done at: -
Status: Not done

done 018f3b5e-...
Task marked as done: 018f3b5e-...

events
[21:14:00] add BuyMilk 2 liters
[21:15:10] done 018f3b5e-...

---

## Project Structure

todo-list/
├── main.go
├── go.mod
├── README.md
├── todo/
│ ├── task.go
│ └── list.go
└── scanner/
├── scanner.go
├── prints.go
├── events.go
└── errors.go

---

## Installation & Run

git clone <your-repo-url>
cd todo-cli
go run main.go

---

## Version History

### v1

- Basic CLI
- Tasks identified by title
- Minimal functionality

### v2 (current)

- UUID-based tasks
- Clean CLI output
- Event logging
- Professional code comments
- Improved structure

---

## Future Improvements (Roadmap)

### v3 - UX Improvements

- Human-friendly task references (`task1`, `task2`, …)
- Automatic mapping: number → UUID
- Cleaner list output

### v4 - Command Map

- Replace switch-case with a command registry
- Cleaner architecture
- Easier to extend

### v5 - Internal Architecture

- Introduce `internal/`
- Separate layers: CLI / core / storage
- Better package structure

### v6 - Persistence

- Save tasks to JSON
- Load tasks on startup
- Optional autosave

### v7 - Testing

- Unit tests for todo package
- Tests for scanner logic
- Mocking input/output

### v8 - HTTP API Version

- Convert todo list into REST API
- Add handlers, routing, JSON responses
- Optional Postman collection

---

## Notes

This project is part of my Go learning journey.  
The repository intentionally shows incremental progress - from simple CLI to more advanced architecture.
