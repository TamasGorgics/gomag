package domain

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ActorName   string
	Version     int64
}

func NewTask(title, description, actorName string) *Task {
	return &Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ActorName:   actorName,
		Version:     1,
	}
}

func (t *Task) MarkAsCompleted() {
	t.flipCompleted(true)
}

func (t *Task) MarkAsIncomplete() {
	t.flipCompleted(false)
}

func (t *Task) flipCompleted(c bool) {
	t.Completed = c
	t.UpdatedAt = time.Now()
	t.Version = t.Version + 1
}
