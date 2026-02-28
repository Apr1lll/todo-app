package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(description string, title string) Task {
	return Task{
		Title:       title,
		Description: description,
		Completed:   false,

		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	CompletedTime := time.Now()

	t.Completed = true
	t.CompletedAt = &CompletedTime
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompletedAt = nil
}
