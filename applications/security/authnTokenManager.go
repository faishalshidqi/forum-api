package security

import (
	"forum-api/commons/sql/database"
	"net/http"
	"time"
)

type AuthnTokenManager interface {
	CreateToken(user database.User, secret string, expiresIn time.Duration) (string, error)
	VerifyToken(tokenString string, secret string) (string, error)
	GetBearerToken(header http.Header) (string, error)
}
