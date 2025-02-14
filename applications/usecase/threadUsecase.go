package usecase

import (
	"context"
	"forum-api/domains"
	"time"
)

type threadUsecase struct {
	taskRepository domains.ThreadRepository
	contextTimeout time.Duration
}

func (tu *threadUsecase) GetById(c context.Context, id string) (*domains.Thread, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.GetById(ctx, id)
}

func (tu *threadUsecase) Add(c context.Context, task domains.AddThreadRequest, owner string) (domains.AddThreadResponseData, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Add(ctx, task, owner)
}

func NewThreadUsecase(taskRepository domains.ThreadRepository, timeout time.Duration) domains.ThreadUsecase {
	return &threadUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}
