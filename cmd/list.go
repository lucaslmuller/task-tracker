package cmd

import (
	"fmt"

	"github.com/lucaslmuller/tasktracker/internal/storage"
	"github.com/lucaslmuller/tasktracker/internal/task"
)

func list(s *storage.Storage, args []string) {
	var tasks []task.Task

	if len(args) == 0 {
		tasks = s.GetTasks()
	} else {
		tasks = s.GetByStatus(args[0])
	}

	fmt.Println("id\tstatus\tdescription")
	for _, t := range tasks {
		fmt.Printf("%d\t%s\t%s\n", t.Id, t.Status, t.Description)
	}
}
