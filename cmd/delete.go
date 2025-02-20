package cmd

import (
	"fmt"
	"strconv"

	"github.com/lucaslmuller/tasktracker/cmd/storage"
)

func delete(s *storage.Storage, args []string) {
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

	s.DeleteTask(id)
	fmt.Printf("Deleted task with id %d\n", id)
}
