package cmd

import (
	"fmt"
	"strconv"

	"github.com/lucaslmuller/tasktracker/cmd/storage"
)

func update(s *storage.Storage, args []string) {
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

	task.Description = args[1]

	s.UpdateTask(*task)
	fmt.Printf("Updated task with id %d\n", id)
}
