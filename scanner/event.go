package scanner

import "time"

// Event represents a single CLI action performed by the user.
// It stores the raw input, the command result, and the timestamp.
type Event struct {
    Input     string    // raw user input (e.g. "add BuyMilk 2 liters")
    Result    string    // command result (empty = success, non-empty = error)
    Timestamp time.Time // when the event occurred
}

// NewEvent creates a new Event instance.
// It is called after each processed command in scanner.go.
func NewEvent(result string, input string) Event {
    return Event{
        Input:     input,
        Result:    result,
        Timestamp: time.Now(),
    }
}
