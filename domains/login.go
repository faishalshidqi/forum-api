package domains

import (
	"context"
	"forum-api/commons/sql/database"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUsecase interface {
	GetUserByUsername(c context.Context, username string) (database.User, error)
	CreateAccessToken(user database.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user database.User, secret string, expiry int) (refreshToken string, err error)
}
