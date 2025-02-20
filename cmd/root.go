package cmd

import (
	"fmt"
	"os"

	"github.com/lucaslmuller/tasktracker/cmd/storage"
)

func Execute(s *storage.Storage) {
	args := os.Args[1:]

	switch args[0] {
	case "add":
		add(s, args[1:])
	case "list":
		list(s, args[1:])
	case "delete":
		delete(s, args[1:])
	case "update":
		update(s, args[1:])
	case "mark-done":
		markDone(s, args[1:])
	case "mark-in-progress":
		markInProgress(s, args[1:])
	default:
		fmt.Println("Invalid command")
	}
}
