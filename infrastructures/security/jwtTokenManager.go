package security

import (
	"errors"
	"fmt"
	"forum-api/applications/security"
	"forum-api/commons/sql/database"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type jwtTokenManager struct {
	authnTokenManager security.AuthnTokenManager
}

func (jtm *jwtTokenManager) GetBearerToken(header http.Header) (string, error) {
	headers := header.Get("Authorization")
	if len(strings.Split(headers, " ")) != 2 {
		return "", errors.New("invalid Authorization header")
	}
	return strings.Split(headers, " ")[1], nil
}

func (jtm *jwtTokenManager) CreateToken(user database.User, secret string, expiresIn time.Duration) (string, error) {
	userID := user.ID
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ForumAPI",
			Subject:   userID.String(),
		},
	)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (jtm *jwtTokenManager) VerifyToken(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return "", err
	}
	subject, err := token.Claims.GetSubject()
	if err != nil {
		return "", err
	}
	expirationTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		return "", err
	}
	issuedAt, err := token.Claims.GetIssuedAt()
	if err != nil {
		return "", err
	}
	claims := &jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   subject,
		ExpiresAt: expirationTime,
		IssuedAt:  issuedAt,
	}
	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}
	if claims.ExpiresAt.Before(time.Now()) {
		return "", errors.New("token expired")
	}
	return claims.Subject, nil
}

func NewJwtTokenManager(authnTokenManager security.AuthnTokenManager) security.AuthnTokenManager {
	return &jwtTokenManager{
		authnTokenManager: authnTokenManager,
	}
}
