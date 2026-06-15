package todo

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// task represents a single ToDo item with metadata such as
// creation time, completion status and optional completion timestamp.
type Task struct {
    ID        uuid.UUID   // Unique identifier for the task
    Title     string      // Task title
    Text      string      // Task description
    IsDone    bool        // Completion status

    CreatedAt time.Time   // Timestamp when task was created
    DoneAt    *time.Time  // Timestamp when task was marked as done (nil if not done)
}

// NewTask creates a new task instance.
// It validates input, generates a UUIDv7 and returns a fully initialized task.
func NewTask(title string, text string) (*Task, error) {
    if title == "" {
        return nil, fmt.Errorf("title is empty")
    }
    if text == "" {
        return nil, fmt.Errorf("text is empty")
    }

    id, err := uuid.NewV7()
    if err != nil {
        return nil, fmt.Errorf("failed to generate UUID v7: %w", err)
    }

    return &Task{
        ID:        id,
        Title:     title,
        Text:      text,
        IsDone:    false,
        CreatedAt: time.Now(),
        DoneAt:    nil,
    }, nil
}

// Done marks the task as completed and sets the completion timestamp.
func (t *Task) Done() {
    doneTime := time.Now()
    t.IsDone = true
    t.DoneAt = &doneTime
}
