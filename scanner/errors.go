package scanner

// These constants represent common command results and error messages.
// They are returned by command handlers in scanner.go.

const (
    needExit       = "exit"                    // special marker used to exit the CLI loop
    emptyInput     = "empty input"             // user pressed Enter without typing anything
    wrongArgs      = "wrong number of arguments" // command received incorrect number of arguments
    unknownCommand = "unknown command"         // command name is not recognized
)
