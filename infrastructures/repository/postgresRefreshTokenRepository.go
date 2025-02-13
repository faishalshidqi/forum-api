package repository

import (
	"context"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
)

type postgresRefreshTokenRepository struct {
	database bootstrap.Database
}

func (p *postgresRefreshTokenRepository) Add(c context.Context, refreshToken string) error {
	_, err := p.database.Query.AddToken(c, refreshToken)
	return err
}

func (p *postgresRefreshTokenRepository) Fetch(c context.Context, refreshToken string) (string, error) {
	token, err := p.database.Query.GetToken(c, refreshToken)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (p *postgresRefreshTokenRepository) Delete(c context.Context, refreshToken string) error {
	_, err := p.database.Query.DeleteToken(c, refreshToken)
	return err
}

func NewPostgresRefreshTokenRepository(database bootstrap.Database) domains.RefreshTokenRepository {
	return &postgresRefreshTokenRepository{
		database: database,
	}
}
