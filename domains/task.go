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

type AddTaskRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type AddTaskResponse struct {
	Message string              `json:"message"`
	Status  string              `json:"status"`
	Data    AddTaskResponseData `json:"data"`
}

type AddTaskResponseData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Owner string `json:"owner"`
}

type TaskRepository interface {
	Add(c context.Context, task Task) (Task, error)
}

type TaskUsecase interface {
	Add(c context.Context, task Task) (Task, error)
}
