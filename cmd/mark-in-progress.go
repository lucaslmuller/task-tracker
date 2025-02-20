package cmd

import (
	"fmt"
	"strconv"

	"github.com/lucaslmuller/tasktracker/cmd/storage"
)

func markInProgress(s *storage.Storage, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid id")
		return
	}

	task := s.GetTaskById(id)
	if task == nil {
		fmt.Println("Task not found")
		return
	}

	task.Status = "in-progress"

	s.UpdateTask(*task)
	fmt.Printf("task with id %d marked as in progress\n", id)
}
