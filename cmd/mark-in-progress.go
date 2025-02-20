package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/lucaslmuller/tasktracker/cmd/storage"
)

func markInProgress(s *storage.Storage, args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("invalid id")
	}

	task, err := s.GetTaskById(id)
	if err != nil {
		return err
	}

	task.Status = "in-progress"

	if err = s.UpdateTask(*task); err != nil {
		return err
	}

	fmt.Printf("task with id %d marked as in progress\n", id)
	return nil
}
