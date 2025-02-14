package usecase

import (
	"context"
	"forum-api/domains"
	"time"
)

type taskUsecase struct {
	taskRepository domains.TaskRepository
	contextTimeout time.Duration
}

func (tu *taskUsecase) Add(ctx context.Context, task domains.Task) (domains.Task, error) {
	//TODO implement me
	panic("implement me")
}

func NewTaskUsecase(taskRepository domains.TaskRepository, timeout time.Duration) domains.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}
