package main

import (
	"log"
	"toDoList/scanner"
	"toDoList/todo"
)

func main() {
    // Create a new todo list (assumes NewList returns *todo.List)
    list := todo.NewList()

    // Create scanner controller and start CLI loop
    s := scanner.NewScanner(list)
    s.Start()

    // If Start returns (e.g., on exit), log and finish
    log.Println("CLI stopped")
}
