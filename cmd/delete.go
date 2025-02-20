package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/lucaslmuller/tasktracker/cmd/storage"
)

func delete(s *storage.Storage, args []string) error {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("invalid id")
	}

	_, err = s.GetTaskById(id)
	if err != nil {
		return err
	}

	if err = s.DeleteTask(id); err != nil {
		return err
	}

	fmt.Printf("Deleted task with id %d\n", id)
	return nil
}
