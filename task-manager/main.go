package main

import (
	"fmt"
	"os"
	// "time"
)

type TaskStatus int

const (
	Todo TaskStatus = iota
	InProgress
	Done
)

type Task struct {
	Name        string
	DueDate     string
	Status      TaskStatus
	Description string
	Recurrent   bool
}

const tasksFiles = "tasks.csv"

func main() {
	readTasksFromFile(tasksFiles)
}

func readTasksFromFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("File content:", string(content))
}
