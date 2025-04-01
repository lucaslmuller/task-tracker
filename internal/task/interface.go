package task

type ITask interface {
	UpdateDescription(description string)
	UpdateStatus(status string)
}
