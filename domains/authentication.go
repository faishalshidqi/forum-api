package domains

import (
	"context"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthenticationResponse struct {
	Status string            `json:"status"`
	Data   AuthnResponseData `json:"data"`
}

type AuthnResponseData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthenticationUsecase interface {
	GetUserByUsername(c context.Context, username string) (User, error)
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user User, secret string, expiry int) (refreshToken string, err error)
	ValidateToken(token, secret string) (string, error)
	CheckPasswordHash(password, hash string) error
}
