package security

import "forum-api/applications/security"

type bcryptPasswordHash struct {
	passwordHash security.PasswordHash
}

func (bph *bcryptPasswordHash) HashPassword(password string) (string, error) {
	hashed, err := bph.passwordHash.HashPassword(password)
	if err != nil {
		return "", err
	}
	return hashed, nil
}

func (bph *bcryptPasswordHash) CheckPasswordHash(password, hash string) error {
	return bph.passwordHash.CheckPasswordHash(password, hash)
}

func NewBcryptPasswordHash(passwordHash security.PasswordHash) security.PasswordHash {
	return &bcryptPasswordHash{
		passwordHash: passwordHash,
	}
}
