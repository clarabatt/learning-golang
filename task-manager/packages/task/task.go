package task

import "fmt"

type TaskStatus int
type Priority int

type Task struct {
	ID          int
	Name        string
	DueDate     string
	Status      TaskStatus
	Description string
	Recurrent   bool
	Priority    Priority
}

const (
	Low Priority = iota
	Medium
	High
)

const (
	Todo TaskStatus = iota
	InProgress
	Done
)

func NewTask(record []string) (*Task, error) {
	return &Task{
		Name:        "mock",
		DueDate:     "mock",
		Status:      Todo,
		Description: "mock",
		Recurrent:   false,
		Priority:    High,
	}, nil
}

func (t *Task) PrintTask() {
	fmt.Printf("[%d] %s - %s \n", t.ID, t.Name, t.DueDate)
}
