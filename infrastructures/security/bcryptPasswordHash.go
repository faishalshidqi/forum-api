package security

import (
	"forum-api/applications/security"
	"golang.org/x/crypto/bcrypt"
)

type bcryptPasswordHash struct{}

func (bph *bcryptPasswordHash) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (bph *bcryptPasswordHash) CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func NewBcryptPasswordHash() security.PasswordHash {
	return &bcryptPasswordHash{}
}
