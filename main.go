package main

import (
	"github.com/lucaslmuller/tasktracker/cmd"
	"github.com/lucaslmuller/tasktracker/cmd/storage"
)

func main() {
	s := storage.Setup()
	cmd.Execute(s)
}
