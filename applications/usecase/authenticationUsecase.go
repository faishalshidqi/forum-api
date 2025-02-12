package usecase

import (
	"context"
	"forum-api/applications/security"
	"forum-api/domains"
	"time"
)

type authenticationUsecase struct {
	userRepository domains.UserRepository
	tokenManager   security.AuthnTokenManager
	passwordHash   security.PasswordHash
	contextTimeout time.Duration
}

func (au *authenticationUsecase) CheckPasswordHash(password, hash string) error {
	return au.passwordHash.CheckPasswordHash(password, hash)
}

func (au *authenticationUsecase) GetUserByID(c context.Context, id string) (domains.User, error) {
	user, err := au.userRepository.GetByID(c, id)
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (au *authenticationUsecase) ValidateToken(token, secret string) (string, error) {
	id, err := au.tokenManager.VerifyToken(token, secret)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (au *authenticationUsecase) GetUserByUsername(c context.Context, username string) (domains.User, error) {
	user, err := au.userRepository.GetByUsername(c, username)
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (au *authenticationUsecase) CreateAccessToken(user domains.User, secret string, expiry int) (accessToken string, err error) {
	accessToken, err = au.tokenManager.CreateToken(user, secret, time.Duration(expiry)*time.Hour)
	return
}

func (au *authenticationUsecase) CreateRefreshToken(user domains.User, secret string, expiry int) (refreshToken string, err error) {
	refreshToken, err = au.tokenManager.CreateToken(user, secret, time.Duration(expiry)*time.Hour)
	return
}

func NewAuthenticationUsecase(userRepository domains.UserRepository, tokenManager security.AuthnTokenManager, passwordHash security.PasswordHash, timeout time.Duration) domains.AuthenticationUsecase {
	return &authenticationUsecase{
		userRepository: userRepository,
		tokenManager:   tokenManager,
		passwordHash:   passwordHash,
		contextTimeout: timeout,
	}
}
