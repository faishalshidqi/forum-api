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

func (tu *taskUsecase) Add(c context.Context, task domains.Task) (domains.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Add(ctx, task)
}

func NewTaskUsecase(taskRepository domains.TaskRepository, timeout time.Duration) domains.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}
