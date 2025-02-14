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
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		Thread   GetThreadByIDResponseData         `json:"thread"`
		Comments []GetCommentsByThreadResponseData `json:"comments"`
	}
}

type GetCommentsByThreadResponseData struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Date     time.Time `json:"date"`
	Content  string    `json:"content"`
}

type CommentRepository interface {
	Add(c context.Context, commentRequest AddCommentRequest, owner, thread string) (AddCommentResponseData, error)
	GetByThread(c context.Context, thread string) ([]Comment, error)
}

type CommentUsecase interface {
	Add(c context.Context, commentRequest AddCommentRequest, owner, thread string) (AddCommentResponseData, error)
	GetByThread(c context.Context, thread string) ([]Comment, error)
}
