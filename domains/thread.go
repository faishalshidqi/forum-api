package domains

import (
	"context"
	"time"
)

type Thread struct {
	ID    string    `json:"id"`
	Title string    `json:"title" binding:"required"`
	Body  string    `json:"body" binding:"required"`
	Date  time.Time `json:"date"`
	Owner string    `json:"owner"`
}

type AddThreadRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type AddThreadResponse struct {
	Message string                `json:"message"`
	Status  string                `json:"status"`
	Data    AddThreadResponseData `json:"data"`
}

type AddThreadResponseData struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Owner string `json:"owner"`
}

type ThreadRepository interface {
	Add(c context.Context, task Thread) (AddThreadResponseData, error)
}

type ThreadUsecase interface {
	Add(c context.Context, task Thread) (AddThreadResponseData, error)
}
