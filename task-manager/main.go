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
	status      TaskStatus
	description string
	recurrent   bool
}

const tasks_files = "tasks.csv"

func main() {

	readingTasks(tasks_files)

}

func readingTasks(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("File content:", string(content))
}
