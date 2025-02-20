package storage

import (
	"encoding/json"
	"os"

	"github.com/lucaslmuller/tasktracker/types"
)

type Storage struct {
	Id    int
	Tasks []types.Task
}

func Setup() *Storage {
	setupDataFile()

	tasks, largestId := readTasks()
	return &Storage{
		Id:    largestId + 1,
		Tasks: tasks,
	}
}

func setupDataFile() {
	os.Mkdir("data", os.ModePerm)
	_, err := os.ReadFile("data/data.json")
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		} else {
			os.WriteFile("data/data.json", []byte("[]"), os.ModePerm)
		}
	}
}

func readTasks() (tasks []types.Task, largestId int) {
	tasks = []types.Task{}
	largestId = 0

	file, err := os.ReadFile("data/data.json")
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	json.Unmarshal(file, &tasks)

	for _, t := range tasks {
		if t.Id > largestId {
			largestId = t.Id
		}
	}

	return tasks, largestId
}
