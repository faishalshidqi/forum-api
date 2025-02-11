package usecase

import (
	"context"
	"forum-api/applications/security"
	"forum-api/commons/sql/database"
	"forum-api/domains"
	"time"
)

type loginUsecase struct {
	userRepository domains.UserRepository
	tokenManager   security.AuthnTokenManager
	contextTimeout time.Duration
}

func (lu loginUsecase) GetUserByUsername(c context.Context, username string) (database.User, error) {
	user, err := lu.userRepository.GetByUsername(c, username)
	if err != nil {
		return database.User{}, err
	}
	return user, nil
}

func (lu loginUsecase) CreateAccessToken(user database.User, secret string, expiry int) (accessToken string, err error) {
	accessToken, err = lu.tokenManager.CreateToken(user, secret, time.Duration(expiry)*time.Hour)
	return
}

func (lu loginUsecase) CreateRefreshToken(user database.User, secret string, expiry int) (refreshToken string, err error) {
	refreshToken, err = lu.tokenManager.CreateToken(user, secret, time.Duration(expiry)*time.Hour)
	return
}

func NewLoginUsecase(userRepository domains.UserRepository, tokenManager security.AuthnTokenManager, timeout time.Duration) domains.LoginUsecase {
	return loginUsecase{
		userRepository: userRepository,
		tokenManager:   tokenManager,
		contextTimeout: timeout,
	}
}
