package domains

import (
	"context"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Status string            `json:"status"`
	Data   LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type RefreshResponse struct {
	Status string              `json:"status"`
	Data   RefreshResponseData `json:"data"`
}

type RefreshResponseData struct {
	AccessToken string `json:"accessToken"`
}

type AuthenticationUsecase interface {
	GetUserByUsername(c context.Context, username string) (User, error)
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user User, secret string, expiry int) (refreshToken string, err error)
	ValidateToken(token, secret string) (string, error)
	CheckPasswordHash(password, hash string) error
}

type RefreshTokenRepository interface {
	Add(c context.Context, refreshToken string) error
	Fetch(c context.Context, refreshToken string) (string, error)
	Delete(c context.Context, refreshToken string) error
}
