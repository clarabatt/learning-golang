package task

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type TaskStatus int
type Priority int

type Task struct {
	ID          int
	Name        string
	DueDate     time.Time
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

func intToTaskStatus(i int) (TaskStatus, error) {
	switch i {
	case 0:
		return Todo, nil
	case 1:
		return InProgress, nil
	case 2:
		return Done, nil
	default:
		return 0, fmt.Errorf("unknown TaskStatus for value: %d", i)
	}
}

func NewTask(record []string) (*Task, error) {
	dateParsed, err := time.Parse("2006-01-02", record[1])
	if err != nil {
		panic(err)
	}

	intStatus, err := strconv.Atoi(record[2])
	if err != nil {
		panic(err)
	}
	status, err := intToTaskStatus(intStatus)

	recurrent, err := strconv.ParseBool(record[4])
	if err != nil {
		panic(err)
	}

	randomId := rand.Intn(900) + 100

	return &Task{
		ID:          randomId,
		Name:        record[0],
		DueDate:     dateParsed,
		Status:      status,
		Description: record[3],
		Recurrent:   recurrent,
		Priority:    High,
	}, nil
}

func (t *Task) PrintTask() {
	fmt.Printf("[%d] %s - %s \n", t.ID, t.Name, t.DueDate)
}
