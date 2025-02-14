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
	Add(ctx context.Context, task Task) (Task, error)
}

type TaskUsecase interface {
	Add(ctx context.Context, task Task) (Task, error)
}
