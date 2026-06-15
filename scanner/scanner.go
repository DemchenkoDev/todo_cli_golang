package scanner

import (
	"bufio"
	"os"
	"strings"
	"toDoList/todo"

	"github.com/google/uuid"
)

// Scanner is the CLI controller.
// It reads user input, parses commands, executes actions on the todo list,
// prints results, and stores command history.
type Scanner struct {
    todoList     *todo.List      // reference to the task list (model)
    bufioScanner *bufio.Scanner  // reads lines from stdin
    events       []Event         // command history
}

// NewScanner creates a new Scanner instance.
// main.go passes an existing todo.List here.
// Scanner does not create or own the list — it only operates on it.
func NewScanner(todoList *todo.List) Scanner {
    return Scanner{
        todoList:     todoList,
        bufioScanner: bufio.NewScanner(os.Stdin),
    }
}

// Start runs the main CLI loop.
// It repeatedly prints a prompt, reads a command, processes it,
// prints errors/help when needed, and logs the event.
func (s *Scanner) Start() {
    for {
        // print ">"
        printPrompt()

        // read a line; if input is closed (Ctrl+D), exit loop
        if !s.bufioScanner.Scan() {
            break
        }

        // raw user input
        inputString := s.bufioScanner.Text()

        // process the command and get a result (empty = success)
        result := s.process(inputString)

        // if result is not empty, it's either an error or "exit"
        if result != "" {

            // exit command
            if result == needExit {
                printExit()
                return
            }

            // print error message
            printResult(result)

            // show help after errors
            printHelp()
        }

        // log the command and its result
        event := NewEvent(result, inputString)
        s.events = append(s.events, event)
    }
}

// process parses the input string, extracts the command,
// and dispatches it to the corresponding handler.
func (s *Scanner) process(inputString string) string {
    fields := strings.Fields(inputString)

    if len(fields) == 0 {
        return emptyInput
    }

    cmd := fields[0]

    switch cmd {
    case "exit":
        return needExit
    case "add":
        return s.cmdAdd(fields)
    case "list":
        return s.cmdList(fields)
    case "del":
        return s.cmdDel(fields)
    case "done":
        return s.cmdDone(fields)
    case "events":
        return s.cmdEvents(fields)
    case "help":
        return s.cmdHelp(fields)
    default:
        return unknownCommand
    }
}

// add <title> <text>
// Creates a new task and adds it to the list.
func (s *Scanner) cmdAdd(fields []string) string {
    if len(fields) < 3 {
        return wrongArgs
    }

    title := fields[1]
    text := strings.Join(fields[2:], " ")

    task, err := todo.NewTask(title, text)
    if err != nil {
        return err.Error()
    }

    s.todoList.AddTask(*task)

    printAdd(task.ID.String(), title)
    return ""
}

// list
// Prints all tasks.
func (s *Scanner) cmdList(fields []string) string {
    if len(fields) != 1 {
        return wrongArgs
    }

    tasks := s.todoList.ListTasks()
    printTasks(tasks)

    return ""
}

// done <UUID>
// Marks a task as completed.
func (s *Scanner) cmdDone(fields []string) string {
    if len(fields) != 2 {
        return wrongArgs
    }

    id, err := uuid.Parse(fields[1])
    if err != nil {
        return "invalid UUID format"
    }

    result := s.todoList.DoneTask(id)
    if result != "" {
        return result
    }

    printDone(id.String())
    return ""
}

// del <UUID>
// Deletes a task from the list.
func (s *Scanner) cmdDel(fields []string) string {
    if len(fields) != 2 {
        return wrongArgs
    }

    id, err := uuid.Parse(fields[1])
    if err != nil {
        return "invalid UUID format"
    }

    result := s.todoList.DeleteTask(id)
    if result != "" {
        return result
    }

    printDel(id.String())
    return ""
}

// events
// Prints command history.
func (s *Scanner) cmdEvents(fields []string) string {
    if len(fields) != 1 {
        return wrongArgs
    }

    printEvents(s.events)
    return ""
}

// help
// Prints available commands.
func (s *Scanner) cmdHelp(fields []string) string {
    if len(fields) != 1 {
        return wrongArgs
    }

    printHelp()
    return ""
}
