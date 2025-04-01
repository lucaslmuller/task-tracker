package storage

import "github.com/lucaslmuller/tasktracker/internal/task"

type IStorage interface {
	AddTask(task task.Task) error
	DeleteTask(id int) error
	UpdateTask(task task.Task) error
	GetTaskById(id int) (*task.Task, error)
	GetByStatus(status string) []task.Task
	GetTasks() []task.Task
	Persist() error
}
