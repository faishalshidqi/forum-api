package security

import (
	"forum-api/domains"
	"time"
)

type AuthnTokenManager interface {
	CreateToken(user domains.User, secret string, expiresIn time.Duration) (string, error)
	VerifyToken(token string, secret string) (string, error)
}
