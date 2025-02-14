package domains

import (
	"context"
	"time"
)

type Task struct {
	ID    string    `json:"id"`
	Title string    `json:"title" binding:"required"`
	Body  string    `json:"body" binding:"required"`
	Date  time.Time `json:"date"`
	Owner string    `json:"owner"`
}

type TaskRepository interface {
	AddTask(ctx context.Context, task Task) (Task, error)
}
