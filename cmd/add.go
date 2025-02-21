package cmd

import (
	"github.com/lucaslmuller/tasktracker/internal/storage"
	"github.com/lucaslmuller/tasktracker/internal/task"
)

func add(s *storage.Storage, args []string) error {
	return s.AddTask(task.NewTask(args[0]))
}
