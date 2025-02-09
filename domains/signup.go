package domains

import "context"

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByUsername(c context.Context, username string) (User, error)
}
