package cmd

import (
	"fmt"
	"os"

	"github.com/lucaslmuller/tasktracker/internal/storage"
)

func Execute(s *storage.Storage) {
	var err error
	args := os.Args[1:]

	switch args[0] {
	case "add":
		err = add(s, args[1:])
	case "list":
		list(s, args[1:])
	case "delete":
		err = delete(s, args[1:])
	case "update":
		err = update(s, args[1:])
	case "mark-done":
		err = markDone(s, args[1:])
	case "mark-in-progress":
		err = markInProgress(s, args[1:])
	default:
		fmt.Println("Invalid command")
	}

	if err != nil {
		fmt.Println(err)
	}
}
