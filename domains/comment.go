package domains

import (
	"context"
	"time"
)

type Comment struct {
	ID        string    `json:"id"`
	Owner     string    `json:"owner"`
	Thread    string    `json:"thread"`
	Content   string    `json:"content"`
	Date      time.Time `json:"date"`
	IsDeleted bool      `json:"is_deleted"`
}

type AddCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

type AddCommentResponse struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    AddCommentResponseData `json:"data"`
}

type AddCommentResponseData struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Owner   string `json:"owner"`
}

type GetCommentsByThreadResponse struct {
	Message string                    `json:"message"`
	Status  string                    `json:"status"`
	Data    GetThreadByIDResponseData `json:"data"`
}

type GetCommentsByThreadResponseData struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Date     time.Time `json:"date"`
	Content  string    `json:"content"`
}

type CommentRepository interface {
	Add(c context.Context, commentRequest AddCommentRequest, owner, thread string) (AddCommentResponseData, error)
	GetByThread(c context.Context, thread string) ([]GetCommentsByThreadResponseData, error)
	GetById(c context.Context, commentId string) (GetCommentsByThreadResponseData, error)
	SoftDelete(c context.Context, id string) error
}

type CommentUsecase interface {
	Add(c context.Context, commentRequest AddCommentRequest, owner, thread string) (AddCommentResponseData, error)
	GetByThread(c context.Context, thread string) ([]GetCommentsByThreadResponseData, error)
	GetById(c context.Context, commentId string) (GetCommentsByThreadResponseData, error)
	SoftDelete(c context.Context, id string) error
}
