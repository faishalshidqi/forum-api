package domains

import (
	"context"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthenticationUsecase interface {
	GetUserByUsername(c context.Context, username string) (User, error)
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user User, secret string, expiry int) (refreshToken string, err error)
	ValidateToken(token, secret string) (string, error)
}
