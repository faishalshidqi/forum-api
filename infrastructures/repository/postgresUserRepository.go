package repository

import (
	"context"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"forum-api/infrastructures/sql/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type postgresUserRepository struct {
	database bootstrap.Database
}

func (pur *postgresUserRepository) Add(ctx context.Context, user domains.SignupRequest) (domains.SignupResponseData, error) {
	returned, err := pur.database.Query.CreateUser(
		ctx,
		database.CreateUserParams{
			Username: user.Username,
			Password: user.Password,
			Fullname: user.Fullname,
		},
	)
	if err != nil {
		return domains.SignupResponseData{}, err
	}
	return returned.ToSignupResponseData(), nil
}

func (pur *postgresUserRepository) Fetch(ctx context.Context) ([]domains.User, error) {
	users, err := pur.database.Query.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	domainsUsers := make([]domains.User, 0)
	for _, user := range users {
		domainsUsers = append(domainsUsers, user.ToDomainsUser())
	}
	return domainsUsers, err
}

func (pur *postgresUserRepository) GetByUsername(ctx context.Context, username string) (domains.User, error) {
	user, err := pur.database.Query.GetByUsername(ctx, username)
	if err != nil {
		return domains.User{}, err
	}
	return user.ToDomainsUser(), nil
}

func (pur *postgresUserRepository) GetByID(ctx context.Context, id string) (domains.User, error) {
	uuid := pgtype.UUID{}
	err := uuid.Scan(id)
	if err != nil {
		return domains.User{}, err
	}
	user, err := pur.database.Query.GetByID(ctx, uuid)
	if err != nil {
		return domains.User{}, err
	}
	return user.ToDomainsUser(), nil
}

func NewPostgresUserRepository(database bootstrap.Database) domains.UserRepository {
	return &postgresUserRepository{
		database: database,
	}
}
