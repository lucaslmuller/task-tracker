package cmd

import (
	"github.com/lucaslmuller/tasktracker/cmd/storage"
	"github.com/lucaslmuller/tasktracker/types"
)

func add(s *storage.Storage, args []string) {
	s.AddTask(types.NewTask(args[0]))
}
