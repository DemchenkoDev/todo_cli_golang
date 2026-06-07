# Go CLI ToDo List

Simple command-line ToDo List application written in Go.

## Features
- Add tasks
- Delete tasks
- Mark tasks as done
- List tasks
- Event logging
- Help menu

## Project Structure
```
.
├── todolist/
│   └── todolist.go
├── main.go
├── go.mod
└── README.md
```

## Run
```
go run main.go
```

## Commands
```
add {title} {text}
delete {title}
done {title}
list
events
help
exit
```
