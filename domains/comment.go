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
	Content string `json:"content"`
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

type CommentRepository interface {
	Add(c context.Context, commentRequest AddCommentRequest, owner string) (AddCommentResponseData, error)
}

type CommentUsecase interface {
	Add(c context.Context, commentRequest AddCommentRequest, owner string) (AddCommentResponseData, error)
}
