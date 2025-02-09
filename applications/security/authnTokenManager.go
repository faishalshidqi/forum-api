package security

import (
	"forum-api/domains"
	"net/http"
	"time"
)

type AuthnTokenManager interface {
	CreateToken(user domains.User, secret string, expiresIn time.Duration) (string, error)
	VerifyToken(tokenString string, secret string) (string, error)
	GetBearerToken(header http.Header) (string, error)
}
