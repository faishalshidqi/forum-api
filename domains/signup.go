package domains

import (
	"context"
	"forum-api/commons/sql/database"
)

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

type SignupResponse struct {
	Message string             `json:"message"`
	Status  string             `json:"status"`
	Data    SignupResponseData `json:"data"`
}

type SignupResponseData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
}

type SignupUsecase interface {
	Create(c context.Context, user *SignupRequest) (SignupResponseData, error)
	GetUserByUsername(c context.Context, username string) (database.User, error)
}
