package types

import "time"

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTask(description string) Task {
	return Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
	t.UpdatedAt = time.Now()
}

func (t *Task) UpdateStatus(status string) {
	t.Status = status
	t.UpdatedAt = time.Now()
}
