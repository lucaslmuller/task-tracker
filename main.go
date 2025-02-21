package main

import (
	"github.com/lucaslmuller/tasktracker/cmd"
	"github.com/lucaslmuller/tasktracker/internal/storage"
)

func main() {
	s := storage.Setup()
	cmd.Execute(s)
}
