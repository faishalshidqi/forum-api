package usecase

import (
	"context"
	"forum-api/commons/sql/database"
	"forum-api/domains"
	"time"
)

type signupUsecase struct {
	userRepository domains.UserRepository
	contextTimeout time.Duration
}

func (su *signupUsecase) Create(c context.Context, user *domains.SignupRequest) (domains.SignupResponseData, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Add(ctx, *user)
}

func (su *signupUsecase) GetUserByUsername(c context.Context, username string) (database.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByUsername(ctx, username)
}

func NewSignupUsecase(userRepository domains.UserRepository, timeout time.Duration) domains.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
