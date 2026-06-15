package todo

import "github.com/google/uuid"

// list represents a collection of tasks stored in a map,
// where each task is indexed by its unique UUID.
// This structure provides operations for adding, listing,
// completing and deleting tasks.
type List struct {
    tasks map[uuid.UUID]Task
}

// NewList creates and returns a new list instance.
// It initializes the internal map that will hold tasks.
// The returned pointer ensures that the list can be modified
// by methods that operate on it.
func NewList() *List {
    return &List{
        tasks: make(map[uuid.UUID]Task),
    }
}

// AddTask inserts a new task into the list.
// The task is stored under its unique ID.
// If a task with the same ID already exists (unlikely with UUID),
// it will be overwritten.
func (l *List) AddTask(t Task) {
    l.tasks[t.ID] = t
}

// ListTasks returns the internal map of tasks.
// The caller receives a reference to the map, not a copy,
// so modifications to the returned map will affect the list.
func (l *List) ListTasks() map[uuid.UUID]Task {
    return l.tasks
}

// DoneTask marks a task as completed by its ID.
// If the task does not exist, it returns taskNotFound.
// On success, it returns an empty string.
// (Later this will be replaced with `error`.)
func (l *List) DoneTask(id uuid.UUID) string {
    t, ok := l.tasks[id]
    if !ok {
        return taskNotFound
    }

    t.Done()
    l.tasks[id] = t

    return ""
}

// DeleteTask removes a task from the list by its ID.
// If the task does not exist, it returns taskNotFound.
// On success, it returns an empty string.
func (l *List) DeleteTask(id uuid.UUID) string {
    _, ok := l.tasks[id]
    if !ok {
        return taskNotFound
    }

    delete(l.tasks, id)
    return ""
}
