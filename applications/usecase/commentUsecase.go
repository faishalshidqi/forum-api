package usecase

import (
	"context"
	"forum-api/domains"
	"time"
)

type commentUsecase struct {
	commentRepository domains.CommentRepository
	timeout           time.Duration
}

func (cu *commentUsecase) Add(c context.Context, commentRequest domains.AddCommentRequest, owner, thread string) (domains.AddCommentResponseData, error) {
	return cu.commentRepository.Add(c, commentRequest, owner, thread)
}

func NewCommentUsecase(commentRepository domains.CommentRepository, timeout time.Duration) domains.CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
		timeout:           timeout,
	}
}
