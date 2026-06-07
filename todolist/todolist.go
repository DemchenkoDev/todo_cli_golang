package application

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

// task represents a single ToDo item.
// Fields are unexported because tasks are managed only inside this package.
type task struct {
	title		string
	text		string
	CreatedAT	time.Time // Timestamp when task was created
	done 		bool	  // Marks whether task is completed
	DoneAT		time.Time // Timestamp when task was marked as done
}

// event stores information about user actions.
// This allows us to keep a history of all commands and errors.
type event struct {
	Input     string    // Raw user input
    Error     string    // Error message if command failed
    CreatedAT time.Time // When the event occurred
}

// NewTask is a constructor for task.
// It validates input and ensures that task always has required fields.
func NewTask(
	title		string,
	text		string,
) (*task, error){
	if title == "" {
		return nil, fmt.Errorf("Title is empty!")
	}
	if text == "" {
		return nil, fmt.Errorf("Text is empty!")
	}
	return &task{
		title: title,
		text: text,
		CreatedAT: time.Now(),
		done: false,
		DoneAT: time.Time{},
	}, nil
	
}

// ToDoList is the main CLI loop.
// It reads user input, parses commands, and updates internal state.
func ToDoList() {
	scanner := bufio.NewScanner(os.Stdin)

	// Preallocate slices to avoid unnecessary reallocations.
	taskSlice := make([]task, 0, 100)
	events := make([]event, 0, 100)

	for {
		fmt.Printf("Please, input the command: ")
		scanner.Scan()
		input := scanner.Text()

		// Temporary slice for delete operation
		var newTasks []task

		if len(input) == 0 {
			fmt.Println("empty input!")
			continue
		}

		words := strings.Fields(input)
		cmd := words[0]

		switch cmd {

		// ---------------- ADD ----------------
		case "add":
			// Create event for logging user action
			e := event{
				Input: input,
				Error: "",
				CreatedAT: time.Now(),
			}

			if len(words) < 3 {
				fmt.Println("Format: add {title} {text}")
				continue
			}

			title := words[1]
			text := strings.Join(words[2:], " ")

			t, err := NewTask(title, text)
				if err != nil {
					e.Error = err.Error()
					fmt.Println("Error: ", err)
					continue
				}

			events = append(events, e)
			taskSlice = append(taskSlice, *t)
			fmt.Println("Task added.")

		// ---------------- LIST ----------------
		case "list":
			// Log event
			e := event{
				Input: input,
				Error: "",
				CreatedAT: time.Now(),
			}

			events = append(events, e)

			// Pretty-print tasks using pp
			pp.Println(taskSlice)

		// ---------------- DELETE ----------------
		case "delete":
			e := event{
				Input: input,
				Error: "",
				CreatedAT: time.Now(),
			}

			title := words[1]

			// Filter out the task with matching title
			for i := 0; i < len(taskSlice); i++ {
				if taskSlice[i].title == title{
					continue
				}
				newTasks = append(newTasks, taskSlice[i])
			}

			events = append(events, e)
			taskSlice = newTasks
			fmt.Println("Deleted task: ", title)

		// ---------------- DONE ----------------
		case "done":
			e := event{
				Input: input,
				Error: "",
				CreatedAT: time.Now(),
			}

			title := words[1]

			// Mark task as done without modifying title
			for i := 0; i < len(taskSlice); i++ {
				if taskSlice[i].title == title && taskSlice[i].done == false{
					taskSlice[i].done = true
					taskSlice[i].DoneAT = time.Now()
				}
			}

			events = append(events, e)
			fmt.Println("Done the task: ", title)

		// ---------------- EVENTS ----------------
		case "events":
			e := event{
				Input: input,
				Error: "",
				CreatedAT: time.Now(),
			}

			// Print all logged events
    		for _, ev := range events {
        		fmt.Printf("[%s] input: %s | error: %s\n",
            	ev.CreatedAT.Format(time.RFC822),
            	ev.Input,
            	ev.Error,
        		)
    		}
			events = append(events, e)

		// ---------------- HELP ----------------
		case "help":
			e := event{
				Input: input,
				Error: "",
				CreatedAT: time.Now(),
			}

			// Simple help menu
			fmt.Println("help: command shows you command list")
			fmt.Println("--------------------------------------------------------------------")
			fmt.Println("add: command allows you add new task: add {title} {text}")
			fmt.Println("--------------------------------------------------------------------")
			fmt.Println("delete: command allows you delete your task: delete {title}")
			fmt.Println("--------------------------------------------------------------------")
			fmt.Println("list: command shows you list of tasks")
			fmt.Println("--------------------------------------------------------------------")
			fmt.Println("done: command allows done your task: done {title}")
			fmt.Println("--------------------------------------------------------------------")
			fmt.Println("events: command shows you list of all events per all timeloop of program life")
			fmt.Println("--------------------------------------------------------------------")
			fmt.Println("exit: command allows exit from programm")
			fmt.Println("--------------------------------------------------------------------")
			events = append(events, e)

		// ---------------- EXIT ----------------
		case "exit":
			// return exits the entire function (not just switch)
			fmt.Println("Exit from program...")
			return

		// ---------------- UNKNOWN ----------------
		default:
			e := event{
				Input: input,
				Error: "",
				CreatedAT: time.Now(),
			}
			fmt.Println("You input the unknown command")
			e.Error = "You input the unknown command"
			events = append(events, e)
		}
	}
}
