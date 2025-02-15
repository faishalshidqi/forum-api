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

func (cu *commentUsecase) GetById(c context.Context, commentId string) (domains.GetCommentsByThreadResponseData, error) {
	return cu.commentRepository.GetById(c, commentId)
}

func (cu *commentUsecase) SoftDelete(c context.Context, id string) error {
	return cu.commentRepository.SoftDelete(c, id)
}

func (cu *commentUsecase) GetByThread(c context.Context, thread string) ([]domains.GetCommentsByThreadResponseData, error) {
	return cu.commentRepository.GetByThread(c, thread)
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
