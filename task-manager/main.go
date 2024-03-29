package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"task-manager/packages/task"
)

const tasksFiles = "files/tasks.csv"

func main() {
	tasksArray := []task.Task{}
	readTasksFromFile(tasksFiles, &tasksArray)

	for _, task := range tasksArray {
		task.PrintTask()
	}
}

func readTasksFromFile(filename string, arr *[]task.Task) {
	content, err := os.ReadFile(filename)
	fmt.Println("Opening the " + filename + " file...")
	if err != nil {
		fmt.Println("Error opening the file...")
		panic(err)
	}

	r := csv.NewReader(strings.NewReader(string(content)))

	_, err = r.Read()
	if err != nil {
		panic(err)
	}

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		newTask, err := task.NewTask(line)
		if err != nil {
			panic(err)
		}
		*arr = append(*arr, *newTask)
	}
}
