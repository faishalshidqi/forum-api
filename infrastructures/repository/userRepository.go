package repository

import (
	"context"
	"forum-api/commons/bootstrap"
	"forum-api/domains"
	"forum-api/infrastructures/sql/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type userRepository struct {
	database bootstrap.Database
}

func (ur *userRepository) Add(ctx context.Context, user domains.SignupRequest) (domains.SignupResponseData, error) {
	returned, err := ur.database.Query.CreateUser(
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
	returnedData := domains.SignupResponseData{}
	returnedData.ID = returned.ID.String()
	returnedData.Username = user.Username
	returnedData.FullName = user.Fullname
	return returnedData, nil
}

func (ur *userRepository) Fetch(ctx context.Context) ([]domains.User, error) {
	users, err := ur.database.Query.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	domainsUsers := make([]domains.User, 0)
	for _, user := range users {
		domainsUsers = append(domainsUsers, user.ToDomainsUser())
	}
	return domainsUsers, err
}

func (ur *userRepository) GetByUsername(ctx context.Context, username string) (domains.User, error) {
	user, err := ur.database.Query.GetByUsername(ctx, username)
	if err != nil {
		return domains.User{}, err
	}
	return user.ToDomainsUser(), nil
}

func (ur *userRepository) GetByID(ctx context.Context, id string) (domains.User, error) {
	uuid := pgtype.UUID{}
	err := uuid.Scan(id)
	if err != nil {
		return domains.User{}, err
	}
	user, err := ur.database.Query.GetByID(ctx, uuid)
	if err != nil {
		return domains.User{}, err
	}
	return user.ToDomainsUser(), nil
}

func NewUserRepository(database bootstrap.Database) domains.UserRepository {
	return &userRepository{
		database: database,
	}
}
