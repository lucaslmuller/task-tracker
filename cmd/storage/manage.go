package storage

import "github.com/lucaslmuller/tasktracker/types"

func (s *Storage) AddTask(task types.Task) {
	task.Id = s.Id
	s.Id++
	s.Tasks = append(s.Tasks, task)

	s.Persist()
}

func (s *Storage) DeleteTask(id int) {
	for i, t := range s.Tasks {
		if t.Id == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
		}
	}

	s.Persist()
}

func (s *Storage) UpdateTask(task types.Task) {
	for i, t := range s.Tasks {
		if t.Id == task.Id {
			s.Tasks[i] = task
		}
	}

	s.Persist()
}

func (s *Storage) GetTaskById(id int) *types.Task {
	for _, t := range s.Tasks {
		if t.Id == id {
			return &t
		}
	}
	return nil
}

func (s *Storage) GetByStatus(status string) []types.Task {
	var tasks []types.Task
	for _, t := range s.Tasks {
		if t.Status == status {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func (s *Storage) GetTasks() []types.Task {
	return s.Tasks
}
