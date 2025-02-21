package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/lucaslmuller/tasktracker/internal/storage"
)

func update(s *storage.Storage, args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("invalid id")
	}

	task, err := s.GetTaskById(id)
	if err != nil {
		return err
	}

	task.Description = args[1]

	if err = s.UpdateTask(*task); err != nil {
		return err
	}

	fmt.Printf("Updated task with id %d\n", id)

	return nil
}
